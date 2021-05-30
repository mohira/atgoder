package topic_integer_problem_prime_factorization

import (
	"atgoder/lib"
	"math"
	"testing"
)

// [ABC057C - Digits in Multiplication](https://atcoder.jp/contests/abc057/tasks/abc057_c)
func AnswerABC057Cその2(N int) int {
	// もっとスマートにいける
	// 約数スライスを2重ループではなく、約数AをみつけたらそのままBも作ればいい
	var ans = math.MaxInt64

	for A := 1; A*A <= N; A++ {
		if N%A == 0 {
			B := N / A

			F := lib.Max(lib.GetDigitNums(A), lib.GetDigitNums(B))

			ans = lib.Min(ans, F)
		}
	}

	return ans
}

func TestAnswerABC057Cその2(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{"入力例1", 10000, 3},
		{"入力例2", 1000003, 7},
		{"入力例3", 9876543210, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC057Cその2(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
