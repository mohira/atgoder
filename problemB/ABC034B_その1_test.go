package problemB

import (
	"testing"
)

// [ABC034B - ペア](https://atcoder.jp/contests/abc034/tasks/abc034_b)
func AnswerABC034Bその1(n int) int {
	if n%2 == 0 {
		return n - 1
	} else {
		return n + 1
	}
}

func TestAnswerABC034Bその1(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"入力例1", 100, 99},
		{"入力例2", 123456789, 123456790},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC034Bその1(tt.n)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
