package topic_amari

import (
	"testing"
)

// [ABC200 - Ringo's Favorite Numbers 2](https://atcoder.jp/contests/abc200/tasks/abc200_c)
func AnswerABC200Cその1(N int, A []int) int {
	const X = 200
	for i := 0; i < N; i++ {
		A[i] %= X
	}

	bucket := make(map[int]int)
	for _, a := range A {
		bucket[a]++
	}

	var count int
	for _, freq := range bucket {
		if freq >= 2 {
			count += nC2(freq)
		}
	}

	return count
}

func nC2(n int) int {
	return n * (n - 1) / 2
}

func TestAnswerABC200Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want int
	}{
		{"入力例1", 6, []int{123, 223, 123, 523, 200, 2000}, 4},
		{"入力例2", 5, []int{1, 2, 3, 4, 5}, 0},
		{"入力例3", 8, []int{199, 100, 200, 400, 300, 500, 600, 200}, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC200Cその1(tt.N, tt.A)
			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
