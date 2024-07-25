package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

type Computer struct {
	CPU string
	RAM int
	MB  string
}

type ComputerBuilderI interface {
	CPU(val string) ComputerBuilderI
	RAM(val int) ComputerBuilderI
	MB(val string) ComputerBuilderI

	Build() Computer
}

type computerBuilder struct {
	cpu string
	ram int
	mb  string
}

func (b computerBuilder) CPU(val string) ComputerBuilderI {
	b.cpu = val
	return b
}
func (b computerBuilder) RAM(val int) ComputerBuilderI {
	b.ram = val
	return b
}
func (b computerBuilder) MB(val string) ComputerBuilderI {
	b.mb = val
	return b
}

func (b computerBuilder) Build() Computer {
	return Computer{
		CPU: b.cpu,
		RAM: b.ram,
		MB:  b.mb,
	}
}
