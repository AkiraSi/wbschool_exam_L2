package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func grep(input []string, pattern string, countFlag, fixedFlag, ignoreCaseFlag, invertFlag, lineNumberFlag bool, afterFlag, beforeFlag, contextFlag int) []string {
	if contextFlag > beforeFlag { // выравнивание a,b в c
		beforeFlag = contextFlag
	}
	if contextFlag > afterFlag {
		afterFlag = contextFlag
	}
	if ignoreCaseFlag { // если регистр не важен
		pattern = strings.ToLower(pattern)
	}
	r, err := regexp.Compile(pattern) // regex
	if err != nil {
		log.Fatalln("Regex:", err)
	}
	var count int = 0 // счетчик совпадений
	var res []string
	for lineNumber, line := range input {
		if ignoreCaseFlag {
			line = strings.ToLower(line) // Если игнорируем регистр, то приводим к нижнему
		}
		match := false
		if fixedFlag && line == pattern { // Если нужно точное совпадение со строкой
			match = true
		} else if r.MatchString(line) {
			match = true
		}
		if invertFlag { // Если нужно найти несовпадения
			match = !match
		}
		if match { // Если совпадение найдено, то выводим его при необходимости
			count++
			if !countFlag { // Если не нужно вывести только кол-во совпадений
				start := lineNumber - beforeFlag // Стартовый индекс в слайсе для вывода строк перед найденной
				if start < 0 {
					start = 0
				}
				for start < lineNumber { // строки до найденнрй
					res = append(res, input[start])
					start++
				}
				if lineNumberFlag { // Выводим номер строки, если задан флаг
					res[lineNumber] += " " + strconv.Itoa(lineNumber+1)
				}
				res = append(res, line)
				end := lineNumber + afterFlag // Индекс последней строки, которую надо вывести после заданной
				if end >= len(input) {
					end = len(input) - 1
				}
				for i := lineNumber + 1; i < end; i++ { // Вывод строк после найденной
					res = append(res, input[i])
				}
			}
		}
	}
	if countFlag {
		res = append(res, strconv.Itoa(count))
	}
	return res
}

func main() {
	var countFlag, fixedFlag, ignoreCaseFlag, invertFlag, lineNumberFlag bool
	var afterFlag, beforeFlag, contextFlag int
	flag.IntVar(&afterFlag, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&beforeFlag, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&contextFlag, "C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&countFlag, "c", false, "количество строк")
	flag.BoolVar(&ignoreCaseFlag, "i", false, "игнорировать регистр")
	flag.BoolVar(&invertFlag, "v", false, "вместо совпадения, исключать")
	flag.BoolVar(&fixedFlag, "F", false, "точное совпадение со строкой, не паттерн")
	flag.BoolVar(&lineNumberFlag, "n", false, "печатать номер строки")
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatalln("Меньше двух аргументов.")
	}
	file, err := os.Open(flag.Arg(1))
	if err != nil {
		log.Fatalln("File:", err)
	}
	defer file.Close() // открытие файла
	scan := bufio.NewScanner(file)
	input := make([]string, 0, 4)
	for scan.Scan() {
		input = append(input, scan.Text()) // запись построчно в слайс строк
	}
	find := flag.Arg(0)
	res := grep(input, find, countFlag, fixedFlag, ignoreCaseFlag, invertFlag, lineNumberFlag, afterFlag, beforeFlag, contextFlag)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}
