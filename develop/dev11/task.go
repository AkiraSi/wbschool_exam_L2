package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

type Events struct {
	events map[int]Event
}

type Event struct {
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Date   string `json:"date"`
	ID     int    `json:"id"`
}

type Response struct {
	Good string `json:"result"`
	Bad  string `json:"error"`
}

func jsonResponse(w http.ResponseWriter, code int, msg interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(msg)
}

func CheckTimeEvent(e Event) bool {
	_, err := time.Parse("2006-01-02", e.Date)
	if err != nil {
		return false
	}
	return true
}

var events Events

func main() {
	http.HandleFunc("/create_event", create_event)         // post
	http.HandleFunc("/update_event", update_event)         // post
	http.HandleFunc("/delete_event", delete_event)         // post
	http.HandleFunc("/events_for_day", events_for_day)     // get
	http.HandleFunc("/events_for_week", events_for_week)   // get
	http.HandleFunc("/events_for_month", events_for_month) // get
	events.events = make(map[int]Event)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func logMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		fmt.Println(r.Method, r.URL.Path)
		f(w, r)
		endTime := time.Now()
		elapsedTime := endTime.Sub(startTime)
		fmt.Println(r.URL.Path, elapsedTime)
	}
}

func create_event(w http.ResponseWriter, r *http.Request) { // post
	if r.Method != http.MethodPost {
		errResp := Response{"", "Метод запрещен"}
		jsonResponse(w, http.StatusMethodNotAllowed, errResp)
		return
	}
	var event Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !CheckTimeEvent(event) {
		http.Error(w, "Неверный формат времени. [2006-01-02]", http.StatusInternalServerError)
		return
	}
	events.events[event.ID] = event
	jsonResponse(w, http.StatusOK, Response{"Good", ""})
}

func update_event(w http.ResponseWriter, r *http.Request) { // post
	if r.Method != http.MethodPost {
		errResp := Response{"", "Метод запрещен"}
		jsonResponse(w, http.StatusMethodNotAllowed, errResp)
		return
	}
	var event Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !CheckTimeEvent(event) {
		http.Error(w, "Неверный формат времени. [2006-01-02]", http.StatusInternalServerError)
		return
	}
	_, find := events.events[event.ID]
	if find {
		events.events[event.ID] = event
		jsonResponse(w, http.StatusOK, Response{"Good", ""})
	} else {
		jsonResponse(w, http.StatusOK, Response{"", "Not found event"})
	}
}

func delete_event(w http.ResponseWriter, r *http.Request) { // post
	fmt.Println(events)
	if r.Method != http.MethodPost {
		errResp := Response{"", "Метод запрещен"}
		jsonResponse(w, http.StatusMethodNotAllowed, errResp)
		return
	}
	var event Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !CheckTimeEvent(event) {
		http.Error(w, "Неверный формат времени. [2006-01-02]", http.StatusInternalServerError)
		return
	}
	_, find := events.events[event.ID]
	if find {
		delete(events.events, event.ID)
		jsonResponse(w, http.StatusOK, Response{"Good", ""})
	} else {
		jsonResponse(w, http.StatusOK, Response{"", "Not found event"})
	}
}

func events_for_day(w http.ResponseWriter, r *http.Request) { // get
	if r.Method != http.MethodGet {
		errResp := Response{"", "Метод запрещен"}
		jsonResponse(w, http.StatusMethodNotAllowed, errResp)
		return
	}
	date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, Response{"", "Bad date"})
		return
	}
	res := make([]Event, 0, 0)
	for _, event := range events.events {
		eventDate, _ := time.Parse("2006-01-02", event.Date)
		if eventDate.Year() == date.Year() && eventDate.Month() == date.Month() && eventDate.Day() == date.Day() {
			res = append(res, event)
		}
	}
	if len(res) < 1 {
		jsonResponse(w, http.StatusBadRequest, Response{"Good, but not events", ""})
		return
	}
	jsonEvents, err := json.Marshal(res)
	jsonResponse(w, http.StatusOK, Response{string(jsonEvents), ""})
}

func events_for_week(w http.ResponseWriter, r *http.Request) { // get
	if r.Method != http.MethodGet {
		errResp := Response{"", "Метод запрещен"}
		jsonResponse(w, http.StatusMethodNotAllowed, errResp)
		return
	}
	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, Response{"", "Bad date"})
		return
	}
	res := make([]Event, 0, 0)
	weekStart := date.AddDate(0, 0, -1)
	weekEnd := weekStart.AddDate(0, 0, 7)
	for _, event := range events.events {
		eventDate, _ := time.Parse("2006-01-02", event.Date)
		if eventDate.After(weekStart) && eventDate.Before(weekEnd) {
			res = append(res, event)
		}
	}
	if len(res) < 1 {
		jsonResponse(w, http.StatusBadRequest, Response{"Good, but not events", ""})
		return
	}
	jsonEvents, err := json.Marshal(res)
	jsonResponse(w, http.StatusOK, Response{string(jsonEvents), ""})
}

func events_for_month(w http.ResponseWriter, r *http.Request) { // get
	if r.Method != http.MethodGet {
		errResp := Response{"", "Метод запрещен"}
		jsonResponse(w, http.StatusMethodNotAllowed, errResp)
		return
	}
	date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, Response{"", "Bad date"})
		return
	}
	res := make([]Event, 0, 0)
	for _, event := range events.events {
		eventDate, _ := time.Parse("2006-01-02", event.Date)
		if eventDate.Year() == date.Year() && eventDate.Month() == date.Month() {
			res = append(res, event)
		}
	}
	if len(res) < 1 {
		jsonResponse(w, http.StatusBadRequest, Response{"Good, but not events", ""})
		return
	}
	jsonEvents, err := json.Marshal(res)
	jsonResponse(w, http.StatusOK, Response{string(jsonEvents), ""})
}
