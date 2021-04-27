package similar

import (
	"strings"
	"testing"
)

// [ABC095A - Something on It](https://atcoder.jp/contests/abc095/tasks/abc095_a)
func AnswerABC095A(S string) int {
	basePrice := 700
	toppingsPrice := 100 * strings.Count(S, "o")

	return basePrice + toppingsPrice
}

func TestABC095A(t *testing.T) {
	tests := []struct {
		name string
		S    string
		want int
	}{
		{"入力例1", "oxo", 900},
		{"入力例2", "ooo", 1000},
		{"入力例3", "xxx", 700},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC095A(tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
