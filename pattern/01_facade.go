package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

type Database interface {
	Save(data string) error
	Load() (string, error)
}

type API interface { // Интерфейс для взаимодействия с API
	SendRequest(url string) (string, error)
}

type MyDatabase struct{} // Реализация базы данных (для примера)

func (db *MyDatabase) Save(data string) error {
	fmt.Println("Сохраняем данные в базу данных:", data)
	return nil
}

func (db *MyDatabase) Load() (string, error) {
	fmt.Println("Загружаем данные из базы данных")
	return "Данные из базы", nil
}

type MyAPI struct{}

func (api *MyAPI) SendRequest(url string) (string, error) {
	fmt.Println("Отправляем запрос в API:", url)
	return "Ответ от API", nil
}

type Architecture struct { // Архитектура
	db  Database
	api API
}

func NewArchitecture(db Database, api API) *Architecture {
	return &Architecture{db: db, api: api}
}

func (sf *Architecture) SaveData(data string) error { // Метод для сохранения данных
	return sf.db.Save(data)
}

func (sf *Architecture) GetDataFromAPI(url string) (string, error) { // Метод для получения данных из API
	return sf.api.SendRequest(url)
}
