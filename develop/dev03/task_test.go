package main

import (
	"reflect"
	"testing"
)

func Test_Sort(t *testing.T) {
	tests := []struct {
		u, n, r bool
		k       int
		input   []string
		want    []string
		err     error
	}{
		{false, false, false, -1, []string{"123 Andrew 1", "1235 Brab 23", "12 Arab 2", "12 Ab", "g dfsg", "g sdf", "12 Ab"}, []string{"12 Ab", "12 Ab", "12 Arab 2", "123 Andrew 1", "1235 Brab 23", "g dfsg", "g sdf"}, nil}, // без параметров
		{true, false, false, -1, []string{"123 Andrew 1", "1235 Brab 23", "12 Arab 2", "12 Ab", "g dfsg", "g sdf", "12 Ab"}, []string{"12 Ab", "12 Arab 2", "123 Andrew 1", "1235 Brab 23", "g dfsg", "g sdf"}, nil},           // unique
		{false, true, false, -1, []string{"123 Andrew 1", "1235 Brab 23", "12 Arab 2", "12 Ab", "g dfsg", "g sdf", "12 Ab"}, []string{"g dfsg", "g sdf", "12 Ab", "12 Ab", "12 Arab 2", "123 Andrew 1", "1235 Brab 23"}, nil},  // numerlic
		{false, false, true, -1, []string{"123 Andrew 1", "1235 Brab 23", "12 Arab 2", "12 Ab", "g dfsg", "g sdf", "12 Ab"}, []string{"g sdf", "g dfsg", "1235 Brab 23", "123 Andrew 1", "12 Arab 2", "12 Ab", "12 Ab"}, nil},  // reverse
		{false, false, false, 1, []string{"123 Andrew 1", "1235 Brab 23", "12 Arab 2", "12 Ab", "g dfsg", "g sdf", "12 Ab"}, []string{"12 Ab", "12 Ab", "12 Arab 2", "123 Andrew 1", "1235 Brab 23", "g dfsg", "g sdf"}, nil},  // key=1
	}
	for _, test := range tests {
		got, err := Sort(test.input, test.n, test.r, test.u, test.k)
		if err == nil && test.err != nil {
			t.Errorf("Sort() error = nil, wantErr %v", test.err)
		}
		if got != nil && !reflect.DeepEqual(got, test.want) {
			t.Errorf("Sort(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}
