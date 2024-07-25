package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func wget() {
	if flag.Arg(0) == "" {
		log.Fatalln("неправильная ссылка")
	}
	resp, err := http.Get(flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	out, err := os.Create("out.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatalln("Could not copy downloaded data to file: ", err.Error())
	}
}

func main() {
	flag.Parse()
	wget()
}
