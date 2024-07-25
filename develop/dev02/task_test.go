package main

import (
	"fmt"
	"testing"
)

func TestUnpackString(t *testing.T) {
	tests := []struct {
		input string
		want  string
		err   error
	}{
		{`a4bc2d5e`, `aaaabccddddde`, nil}, // первая проверочная строка
		{`abcd`, `abcd`, nil},              // вторая проверочная строка
		{`45`, ``, fmt.Errorf("первый символ строки - число")}, // ошибочная строка (3-я)
		{`a10bb`, `aaaaaaaaaabb`, nil},  // строка с двухзначным числом
		{`a15`, `aaaaaaaaaaaaaaa`, nil}, // строка с двухзначным числом + в конце
		{``, ``, nil},                   // пустая строка
	}

	for _, test := range tests {
		got, err := unpack(test.input)
		if err == nil && test.err != nil { // если ошибка должна быть, но ее нет
			t.Errorf("unpack(%q) error = nil, wantErr %v", test.input, test.err)
		}
		if got != test.want { // проверка на соответствие
			t.Errorf("unpack(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}
