package main

import (
	"bytes"
	"os"
	"reflect"
	"strings"
	"testing"
)

func Test_Shell(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
		err   error
	}{
		{"pwd", "pwd", "C:\\Users\\New\\Documents\\GitHub\\wbschool_exam_L2\\develop\\dev08> ", nil},
		{"echo", "echo hello", "hello", nil},
		// {"kill", "kill 1234", "", nil},
		{"cd", "cd C:\\", "", nil},
	}

	for _, tt := range tests {
		tmp := strings.Split(tt.input, " ")
		var buf bytes.Buffer
		RunCommand(tmp, &buf)
		got := buf.String()
		if tt.err != nil {
			t.Errorf("Shell() error = nil, wantErr %v", tt.err)
		}
		if os.Stdout != nil && !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Shell(%q) = %v, want %v", tmp, got, tt.want)
		}
	}
}
