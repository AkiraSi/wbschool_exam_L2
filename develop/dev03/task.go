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
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func sortStrings(input []string, reverse bool) { // сортировка пузырьком + reverse
	for i := 0; i < len(input)-1; i++ {
		for j := 0; j < len(input)-1-i; j++ {
			if (input[j] > input[j+1]) && !reverse {
				input[j], input[j+1] = input[j+1], input[j]
			} else if !(input[j] > input[j+1]) && reverse {
				input[j+1], input[j] = input[j], input[j+1]
			}
		}
	}
}

func sortNumeric(input []string, key int, reverse bool) {
	key-- // индекс
	values := make([]int, len(input))
	for i, line := range input {
		parts := strings.Split(line, " ")
		value, err := strconv.Atoi(parts[key])
		if err != nil {
			values[i] = 0
		} else {
			values[i] = value
		}
	}
	for i := 0; i < len(input)-1; i++ {
		for j := 0; j < len(input)-i-1; j++ {
			if (values[j] < values[j+1] && reverse) || (values[j] > values[j+1] && !reverse) { // сортировка по числовому значению с учетом реверса и без
				input[j], input[j+1] = input[j+1], input[j]
				values[j], values[j+1] = values[j+1], values[j]
			} else if values[j] == values[j+1] { // доп сортировка, если ключ совпал (чтобы была сортировка и по алфавиту дополнительно)
				if strings.Compare(input[j], input[j+1]) > 0 {
					input[j], input[j+1] = input[j+1], input[j]
				}
			}
		}
	}
}

func sortKey(input []string, key int, reverse bool) {
	key-- // индекс
	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			parts1, parts2 := strings.Split(input[i], " "), strings.Split(input[j], " ")
			if len(parts1) <= key || len(parts2) <= key {
				continue // когда нечего сравнивать
			}
			valueI, err := strconv.Atoi(parts1[key])
			if err != nil {
				valueI = int(parts1[key][0]) - 1000000 // по первому символу слова в ключе строки + -1000000
			}
			valueJ, err := strconv.Atoi(parts2[key])
			if err != nil {
				valueJ = int(parts2[key][0]) - 1000000 // по первому символу слова в ключе строки + -1000000
			}
			if (valueI < valueJ && reverse) || (valueI > valueJ && !reverse) {
				input[i], input[j] = input[j], input[i]
			}
		}
	}
}

func uniqueString(input []string) (result []string) {
	for _, item := range input {
		if !contains(result, item) {
			result = append(result, item)
		}
	}
	return result
}

func contains(slice []string, item string) bool {
	for _, element := range slice {
		if element == item {
			return true
		}
	}
	return false
}

func Sort(input []string, numeric bool, reverse bool, unique bool, key int) (result []string, err error) {
	if unique { // уникальность строк, раньше numeric/reverse, для оптимизации
		result = uniqueString(input) // Применяем unique к исходному массиву
	} else {
		result = input // Используем исходный массив, если unique не установлен
	}
	if key >= 1 {
		sortKey(result, key, reverse)
	}
	if numeric && key == -1 {
		sortNumeric(result, 1, reverse)
	} else {
		sortStrings(result, reverse)
	}
	return result, nil
}

func main() {
	unique := flag.Bool("u", false, "Не выводить повторяющиеся строки")
	numeric := flag.Bool("n", false, "Сортировка по числовому значению")
	reverse := flag.Bool("r", false, "Сортировка в обратном порядке")
	key := flag.Int("k", -1, "Номер колонки для сортировки")
	flag.Parse()
	file, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scan := bufio.NewScanner(file)
	input := make([]string, 0, 4)
	for scan.Scan() {
		input = append(input, scan.Text())
	}
	res, _ := Sort(input, *numeric, *reverse, *unique, *key)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}
