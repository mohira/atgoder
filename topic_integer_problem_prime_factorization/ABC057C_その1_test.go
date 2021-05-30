package topic_integer_problem_prime_factorization

import (
	"atgoder/lib"
	"math"
	"testing"
)

// [ABC057C - Digits in Multiplication](https://atcoder.jp/contests/abc057/tasks/abc057_c)
func AnswerABC057Cその1(N int) int {
	// 約数を重複なく収集する
	var divisors []int
	for i := 1; i*i <= N; i++ {
		if N%i == 0 {
			divisors = append(divisors, i)

			if N/i != i {
				divisors = append(divisors, N/i)
			}
		}
	}

	// N = A * B を満たすペアを全探索して、F(A,B)の最小値を探す
	var ans = math.MaxInt64
	for _, A := range divisors {
		for _, B := range divisors {
			if N == A*B {
				digitNumA := lib.GetDigitNums(A)
				digitNumB := lib.GetDigitNums(B)
				ans = lib.Max(digitNumA, digitNumB)
			}
		}
	}
	return ans
}

func TestAnswerABC057Cその1(t *testing.T) {
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
			got := AnswerABC057Cその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
