package similar

import (
	"testing"
)

// [ABC083B - Some Sums](https://atcoder.jp/contests/abc083/tasks/abc083_b)
func AnswerABC083Bその3(N, A, B int) int {
	// 各桁の和を割り算で求めるパターン
	var ans int

	for i := 1; i <= N; i++ {
		totalDigits := calcTotalDigitsWhileっぽく(i)

		if A <= totalDigits && totalDigits <= B {
			ans += i
		}
	}

	return ans
}

func calcTotalDigitsWhileっぽく(n int) int {
	sum := 0

	for n > 0 {
		sum += n % 10
		n /= 10
	}

	return sum
}

func TestAnswerABC083Bその3(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    int
		B    int
		want int
	}{
		{"入力例1", 20, 2, 5, 84},
		//{"入力例2", 10, 1, 2, 13},
		//{"入力例3", 100, 4, 16, 4554},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC083Bその3(tt.N, tt.A, tt.B)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}

		})

	}
}
