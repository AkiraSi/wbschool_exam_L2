package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

type Command interface {
	Execute()
}

type IncrementCommand struct { // Команда для увеличения счетчика
	Counter *int
}

func (ic *IncrementCommand) Execute() {
	*ic.Counter++
}

type DecrementCommand struct { // Команда для уменьшения счетчика
	Counter *int
}

func (dc *DecrementCommand) Execute() {
	*dc.Counter--
}

type Invoker struct { // Invoker - объект, который вызывает команды
	Commands []Command
}

func (iv *Invoker) AddCommand(command Command) {
	iv.Commands = append(iv.Commands, command)
}

func (iv *Invoker) ExecuteCommands() {
	for _, command := range iv.Commands {
		command.Execute()
	}
}
