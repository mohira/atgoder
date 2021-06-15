package ARC_problemA

import (
	"strings"
	"testing"
)

// [ARC050A - 大文字と小文字](https://atcoder.jp/contests/arc050/tasks/arc050_1)
func AnswerARC050Aその1(C, c string) string {
	if strings.ToLower(C) == c {
		return "Yes"
	} else {
		return "No"
	}
}

func TestAnswerARC050Aその1(t *testing.T) {
	tests := []struct {
		name string
		C, c string
		want string
	}{
		{"入力例1", "A", "a", "Yes"},
		{"入力例2", "B", "c", "No"},
		{"入力例3", "Z", "z", "Yes"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerARC050Aその1(tt.C, tt.c)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
