package main

import (
	"fmt"
)

// func test() (x int) { // 0
// 	defer func() {
// 		x++ // 2
// 	}()
// 	x = 1 // 1
// 	return
// }

// func anotherTest() int {
// 	var x int // nil
// 	defer func() {
// 		// fmt.Println(&x, "defer", x)
// 		x++ // 2
// 		// fmt.Println(&x, "defer", x)
// 	}()
// 	x = 1 // 1
// 	// fmt.Println(&x, "return", x)
// 	return x
// }

func foo(s []int) {
	s = append(s, 4)
}

func main() {
	// a := 3
	// fmt.Println(a)
	// fmt.Println(test())
	// fmt.Println(anotherTest())
	// var b = make([]int, a, 6)
	// var c [a]int
	s := []int{1, 2, 3}
	foo(s)
	fmt.Println(len(s))
	fmt.Println(s)
}
