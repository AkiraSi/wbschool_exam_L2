// package for time
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	ctime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err) // вывод в поток вывода stderr ошибки
		os.Exit(-1)                              // выход с ошибкой -1
	}
	fmt.Println(time.Now()) // текущее время
	fmt.Println(ctime)      // точное время
}
