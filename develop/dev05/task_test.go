package main

import (
	"reflect"
	"testing"
)

func Test_Sort(t *testing.T) {
	tests := []struct {
		input         []string
		pattern       string
		A, B, C       int
		c, i, v, F, n bool
		want          []string
		err           error
	}{
		{[]string{"123 Andrew 1", "1235 Brab 23", "12 Arab 2", "12 Ab", "g dfsg", "g sdf", "12 Ab"}, "1", 0, 0, 0, false, false, false, false, false, []string{"123 Andrew 1", "1235 Brab 23", "12 Arab 2", "12 Ab", "12 Ab"}, nil}, // без параметров, поиск 1
	}
	for _, test := range tests {
		got := grep(test.input, test.pattern, test.c, test.F, test.i, test.v, test.n, test.A, test.B, test.C)
		if test.err != nil {
			t.Errorf("Sort() error = nil, wantErr %v", test.err)
		}
		if got != nil && !reflect.DeepEqual(got, test.want) {
			t.Errorf("grep(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}
