package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	fields := flag.String("f", "", "выбрать поля (колонки)")
	delimiter := flag.String("d", "\t", "использовать другой разделитель")
	separated := flag.Bool("s", false, "только строки с разделителем")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if *separated && !strings.Contains(line, *delimiter) {
			continue
		}
		columns := strings.Split(line, *delimiter)
		selectedFields := make([]string, 0)
		if *fields != "" {
			for _, fieldStr := range strings.Split(*fields, ",") {
				field, err := strconv.Atoi(fieldStr)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Invalid field number: %s\n", fieldStr)
					os.Exit(1)
				}
				if field >= 1 && field <= len(columns) {
					selectedFields = append(selectedFields, columns[field-1])
				}
			}
		} else {
			selectedFields = columns
		}
		fmt.Println(strings.Join(selectedFields, *delimiter))
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}
