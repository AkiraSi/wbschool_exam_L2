package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_anogram(t *testing.T) {
	tests := []struct {
		input []string
		want  map[string][]string
		err   error
	}{
		{[]string{"пятак", "пяТка", "тЯпкА", "лиСток", "слитОк", "Столик"}, map[string][]string{"акптя": {"пятак", "пятка", "тяпка"}, "иклост": {"листок", "слиток", "столик"}}, nil}, // первая проверочная строка
		{[]string{"робот", "топот", "карта", "тарка"}, map[string][]string{"аакрт": {"карта", "тарка"}, "боорт": {"робот"}, "ооптт": {"топот"}}, nil},
		{[]string{}, nil, fmt.Errorf("нет ни единого слова")}, //
	}
	for _, test := range tests {
		got, err := anogram(test.input)
		if err == nil && test.err != nil {
			t.Errorf("anogram() error = nil, wantErr %v", test.err)
		}
		if got != nil && !reflect.DeepEqual(*got, test.want) {
			t.Errorf("anogram(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}
