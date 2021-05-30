package topic_integer_problem_prime_factorization

import (
	"testing"
)

// [ARC017A - 素数、コンテスト、素数](https://atcoder.jp/contests/arc017/tasks/arc017_1)
func AnswerARC017Aその1(N int) string {
	if isPrime(N) {
		return "YES"
	} else {
		return "NO"
	}
}

// O(√N)の素数判定
func isPrime(N int) bool {
	// i <= int(math.Sqrt(float64(N))) だと精度が問題になるかもしれない
	for i := 2; i*i <= N; i++ {
		if N%i == 0 {
			return false
		}
	}
	return true
}

func TestAnswerARC017Aその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want string
	}{
		{"入力例1", 17, "YES"},
		{"入力例2", 18, "NO"},
		{"入力例3", 999983, "YES"},
		{"入力例4", 672263, "NO"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerARC017Aその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
