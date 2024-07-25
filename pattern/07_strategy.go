package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

type Strategy interface { // Интерфейс для стратегии
	Execute(a, b int) int
}

type AdditionStrategy struct{} // Стратегия сложения

func (as *AdditionStrategy) Execute(a, b int) int {
	return a + b
}

type SubtractionStrategy struct{} // Стратегия вычитания

func (ss *SubtractionStrategy) Execute(a, b int) int {
	return a - b
}

type ContextS struct { // Контекст, использующий стратегию
	strategy Strategy
}

func (c *ContextS) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *ContextS) ExecuteStrategy(a, b int) int {
	return c.strategy.Execute(a, b)
}
