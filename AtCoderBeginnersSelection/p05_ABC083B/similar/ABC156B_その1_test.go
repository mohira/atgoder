package similar

import (
	"testing"
)

// [ABC156B - Digits](https://atcoder.jp/contests/abc156/tasks/abc156_b)
func AnswerABC156Bその1(N, K int) int {
	var digitNum int

	for N > 0 {
		N /= K
		digitNum++
	}

	return digitNum
}

func TestAnswerABC156Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		K    int
		want int
	}{
		{"入力例1", 11, 2, 4},
		{"入力例2", 1010101, 10, 7},
		{"入力例3", 314159265, 3, 18},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC156Bその1(tt.N, tt.K)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
