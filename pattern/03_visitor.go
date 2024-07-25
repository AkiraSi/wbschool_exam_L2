package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

type Node interface {
	Accept(visitor Visitor)
}

// Интерфейс для посетителя
type Visitor interface {
	Visit(node *TreeNode)
}

// Узел дерева
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) Accept(visitor Visitor) {
	visitor.Visit(node)
}

// Посетитель для печати значения узла
type PrintVisitor struct{}

func (pv *PrintVisitor) Visit(node *TreeNode) {
	fmt.Println(node.Value)
}

// Посетитель для суммирования значений узлов
type SumVisitor struct {
	Sum int
}

func (sv *SumVisitor) Visit(node *TreeNode) {
	sv.Sum += node.Value
}
