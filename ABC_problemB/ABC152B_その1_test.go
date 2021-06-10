package ABC_problemB

import (
	"strconv"
	"strings"
	"testing"
)

// [ABC152B - Comparing Strings](https://atcoder.jp/contests/abc152/tasks/abc152_b)
func AnswerABC152Bその1(a, b int) string {
	s1 := strings.Repeat(strconv.Itoa(a), b)
	s2 := strings.Repeat(strconv.Itoa(b), a)

	if s1 < s2 {
		return s1
	} else {
		return s2
	}
}

func TestAnswerABC152Bその1(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want string
	}{
		{"入力例1", 4, 3, "3333"},
		{"入力例2", 7, 7, "7777777"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC152Bその1(tt.a, tt.b)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
