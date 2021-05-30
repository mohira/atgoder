package topic_integer_problem_prime_factorization

import (
	"testing"
)

// [ABC026B - 完全数](https://atcoder.jp/contests/abc026/tasks/abc026_b)
func AnswerABC026Bその1(N int) string {
	var total = -N // 約数にN自身も含めているので調整する
	for i := 1; i*i <= N; i++ {
		if N%i == 0 {
			total += i

			// Nが平方数のときにダブってしまうのを防いでいる
			// mapで管理するのもありだと思う
			if N/i != i {
				total += N / i
			}
		}
	}

	switch {
	case N < total:
		return "Abundant"
	case total < N:
		return "Deficient"
	case total == N:
		return "Perfect"
	default:
		return ""
	}
}

func TestAnswerABC026Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want string
	}{
		{"入力例1", 6, "Perfect"},
		{"入力例2", 24, "Abundant"},
		{"入力例3", 27, "Deficient"},
		{"入力例4", 945, "Abundant"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC026Bその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
