package main

import (
	"fmt"
	"log"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// anogram функция для поиска множеств анаграмм по словарю.
func anogram(words []string) (*map[string][]string, error) {
	if len(words) == 0 {
		return nil, fmt.Errorf("нет ни единого слова")
	}
	anagrams := make(map[string][]string)
	for _, word := range words {
		word = strings.ToLower(word)
		sortedWord := sortString(word)
		_, ok := anagrams[sortedWord]
		if ok {
			anagrams[sortedWord] = append(anagrams[sortedWord], word)
		} else {
			anagrams[sortedWord] = []string{word}
		}
	}
	return &anagrams, nil
}

func sortString(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)-1; i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] > runes[j] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
	return string(runes)
}

func main() {
	words := []string{}
	anagrams, err := anogram(words)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(*anagrams)
}
