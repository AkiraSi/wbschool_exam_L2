Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) { // 0
	defer func() {
		x++ // 2
	}()
	x = 1 // 1
	return
}


func anotherTest() int {
	var x int // nil
	defer func() {
		x++ // 2
	}()
	x = 1 // 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```
```
Ответ: 2 1 (Отложенные функции изменяют именованные параметры)

```
