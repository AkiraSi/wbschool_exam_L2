package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

type State interface {
	Handle(context *Context)
}

type OnState struct{} // Состояние "Включено"

func (os *OnState) Handle(context *Context) {
	fmt.Println("Состояние: Включено")
	context.SetState(&OffState{})
}

type OffState struct{} // Состояние "Выключено"

func (os *OffState) Handle(context *Context) {
	fmt.Println("Состояние: Выключено")
	context.SetState(&OnState{})
}

type Context struct { // Контекст, использующий состояние
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) Request() {
	c.state.Handle(c)
}
