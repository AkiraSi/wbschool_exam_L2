package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

type Handler interface {
	Handle(request string) bool
}

type BaseHandler struct { // Базовый обработчик
	Next Handler
}

func (bh *BaseHandler) SetNext(handler Handler) {
	bh.Next = handler
}

type Handler1 struct { // Конкретный обработчик 1
	BaseHandler
}

func (h1 *Handler1) Handle(request string) bool {
	if request == "Request 1" {
		fmt.Println("Обработка запроса 1")
		return true
	}
	return false
}

type Handler2 struct { // Конкретный обработчик 2
	BaseHandler
}

func (h2 *Handler2) Handle(request string) bool {
	if request == "Request 2" {
		fmt.Println("Обработка запроса 2")
		return true
	}
	return false
}

type Handler3 struct { // Конкретный обработчик 3
	BaseHandler
}

func (h3 *Handler3) Handle(request string) bool {
	if request == "Request 3" {
		fmt.Println("Обработка запроса 3")
		return true
	}
	return false
}
