package topic_integer_problem_prime_factorization

import (
	"atgoder/lib"
	"math"
	"testing"
)

// [ABC144C - Walk on Multiplication Table ](https://atcoder.jp/contests/abc144/tasks/abc144_c)
func AnswerABC144Cその1(N int) int {
	// 約数のペアを見つけて、総和をとればいい
	minTotal := math.MaxInt64

	for i := 1; i*i <= N; i++ {
		if N%i == 0 {
			minTotal = lib.Min(minTotal, i+N/i)
		}
	}

	// 初期位置は(1,1)なので、2手分調整する
	return minTotal - 2
}

func TestAnswerABC144Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{"入力例1", 10, 5},
		{"入力例2", 50, 13},
		{"入力例3", 10000000019, 10000000018},
		{"入力例3", 1000000000000, 1999998},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC144Cその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
