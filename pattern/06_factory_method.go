package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

type Product interface {
	GetName() string
}

type Product1 struct { // Конкретный продукт 1
	name string
}

func (p *Product1) GetName() string {
	return p.name
}

type Product2 struct { // Конкретный продукт 2
	name string
}

func (p *Product2) GetName() string {
	return p.name
}

type ProductFactory interface { // Интерфейс для фабрики
	CreateProduct() Product
}

type Product1Factory struct{} // Фабрика для продукта 1

func (pf *Product1Factory) CreateProduct() Product {
	return &Product1{name: "Продукт 1"}
}

// Фабрика для продукта 2
type Product2Factory struct{}

func (pf *Product2Factory) CreateProduct() Product {
	return &Product2{name: "Продукт 2"}
}
