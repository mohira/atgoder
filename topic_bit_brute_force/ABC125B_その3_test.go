package topic_bit_brute_force

import (
	"atgoder/lib"
	"testing"
)

// [ABC125B - Resale](https://atcoder.jp/contests/abc125/tasks/abc125_b)
func AnswerABC125Bその3(N int, V, C []int) int {
	// ビット全探索で解いてみる(N<=20なので間に合う)
	var ans = 0
	for bit := 0; bit < (1 << N); bit++ {
		value := 0
		cost := 0
		for i := 0; i < N; i++ {
			if (bit>>i)&1 == 1 {
				value += V[i]
				cost += C[i]
			}
		}

		ans = lib.Max(ans, value-cost)
	}

	return ans
}

func TestAnswerABC125Bその3(t *testing.T) {
	tests := []struct {
		name string
		N    int
		V    []int
		C    []int
		want int
	}{
		{
			"入力例1",
			3,
			[]int{10, 2, 5},
			[]int{6, 3, 4},
			5,
		},
		{
			"入力例2",
			4,
			[]int{13, 21, 6, 19},
			[]int{11, 30, 6, 15},
			6,
		},
		{
			"入力例3",
			1,
			[]int{1},
			[]int{50},
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC125Bその3(tt.N, tt.V, tt.C)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
