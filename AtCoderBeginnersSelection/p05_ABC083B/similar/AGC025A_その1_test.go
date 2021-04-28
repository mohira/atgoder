package similar

import (
	"testing"
)

// [AGC025A - Digits Sum](https://atcoder.jp/contests/agc025/tasks/agc025_a)
func AnswerAGC025Aその1(N int) int {
	// 愚直に全探索
	var min = 1_000_000_000

	for a := 1; a < N; a++ {
		b := N - a

		sum := digitsSum(a) + digitsSum(b)

		if sum < min {
			min = sum
		}
	}

	return min
}

// 各桁の和を求める
func digitsSum(n int) int {
	var sum int

	for n > 0 {
		sum += n % 10
		n /= 10
	}

	return sum
}

func TestAnswerAGC025Aその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{"入力例1", 15, 6},
		{"入力例2", 100000, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerAGC025Aその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
