package ABC_problemB

import (
	"testing"
)

// [ABC111B - AtCoder Beginner Contest 111](https://atcoder.jp/contests/abc111/tasks/abc111_b)
func AnswerABC111Bその1(N int) int {
	for i := 111; i <= 999; i += 111 {
		if N <= i {
			return i
		}
	}

	return 0
}

func TestAnswerABC111Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{"入力例1: ゾロ目", 111, 111},
		{"入力例2", 112, 222},
		{"入力例3", 750, 777},
		{"", 321, 333},
		{"", 231, 333},
		{"", 239, 333},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC111Bその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
