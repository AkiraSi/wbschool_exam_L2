package main

import (
	"fmt"
	"strconv"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func isInt(c rune) bool {
	return unicode.IsDigit(c)
}

func unpack(input string) (string, error) {
	if len(input) == 0 {
		return "", nil
	}
	runeString := []rune(input)
	if isInt(runeString[0]) {
		return "", fmt.Errorf("первый символ строки - число")
	}
	runeF := make([]rune, 0, 4)
	for i := 0; i < len(runeString); i++ {
		if !isInt(runeString[i]) {
			runeF = append(runeF, runeString[i])
		} else {
			var ist string
			temp := -1
			for g := i; g < len(runeString) && isInt(runeString[g]); g++ {
				temp++
				r := runeString[g]
				ist += string(r)
			}
			count, err := strconv.Atoi(ist)
			if err != nil {
				return "", fmt.Errorf("ошибка с конвертацией в число")
			}
			for k := 0; k < count-1; k++ {
				runeF = append(runeF, runeString[i-1])
			}
			i += temp
		}
	}
	return string(runeF), nil
}

func main() {
	s1, err := unpack("a4bc2d5e")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s1)
	}
	s2, err := unpack("abcd")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s2)
	}
	s3, err := unpack("45")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s3)
	}
	s4, err := unpack("")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s4)
	}
}
