package topic_bit_brute_force

import (
	"testing"
)

// [ABC128C - Switches](https://atcoder.jp/contests/abc128/tasks/abc128_c)
func AnswerABC128Cその2(N, M int, KS [][]int, P []int) int {
	// 点灯パターン全列挙作戦
	patterns := make([][]bool, 1<<N)
	for i := 0; i < len(patterns); i++ {
		patterns[i] = make([]bool, N)
	}

	var ans int

	for bit := 0; bit < (1 << N); bit++ {
		// 点灯パターン生成
		pattern := patterns[bit]

		for i := 0; i < N; i++ {
			if (bit>>i)&1 == 1 {
				pattern[i] = true
			} else {
				pattern[i] = false
			}
		}

		// 各点灯パターンにおいて、
		// それぞれの電球の条件を満たすかを調べる
		can := true
		for i := 0; i < M; i++ {
			switches := KS[i][1:]

			count := 0

			for _, s := range switches {
				if pattern[s-1] {
					count++
				}
			}

			if count%2 != P[i] {
				can = false
			}
		}

		if can {
			ans++
		}
	}

	return ans
}
func TestAnswerABC128Cその2(t *testing.T) {
	tests := []struct {
		name string
		N    int
		M    int
		KS   [][]int
		P    []int
		want int
	}{
		{"入力例1",
			2, 2,
			[][]int{
				{2, 1, 2},
				{1, 2},
				{0, 1},
			},
			[]int{0, 1},
			1,
		},
		{"入力例2",
			2, 3,
			[][]int{
				{2, 1, 2},
				{1, 1},
				{1, 2},
			},
			[]int{0, 0, 1},
			0,
		},
		{"入力例3",
			5, 2,
			[][]int{
				{3, 1, 2, 5},
				{2, 2, 3},
				{1, 0},
			},
			[]int{1, 0},
			8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC128Cその2(tt.N, tt.M, tt.KS, tt.P)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
