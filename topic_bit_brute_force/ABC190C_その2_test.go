package topic_bit_brute_force

import (
	"atgoder/lib"
	"testing"
)

// [ABC190C - Bowls and Dishes](https://atcoder.jp/contests/abc190/tasks/abc190_c)
func AnswerABC190Cその2(N, M int, AB [][2]int, K int, CD [][2]int) int {
	patterns := make([][]int, 1<<K)
	for i := 0; i < len(patterns); i++ {
		patterns[i] = make([]int, N+1)
	}

	var maxCount int

	for bit := 0; bit < (1 << K); bit++ {
		// 皿のパターンをつくる; パターン数は2^K
		pattern := patterns[bit]

		for i := 0; i < K; i++ {
			c, d := CD[i][0], CD[i][1]
			if (bit>>i)&1 == 1 {
				pattern[c]++
			} else {
				pattern[d]++
			}
		}

		// 各パターンにおける条件マッチを数える
		count := 0
		for i := 0; i < M; i++ {
			a, b := AB[i][0], AB[i][1]

			if pattern[a] > 0 && pattern[b] > 0 {
				count++
			}
			maxCount = lib.Max(maxCount, count)
		}
	}

	return maxCount
}

func TestAnswerABC190Cその2(t *testing.T) {
	tests := []struct {
		name string
		N    int
		M    int
		AB   [][2]int
		K    int
		CD   [][2]int
		want int
	}{
		{"入力例1",
			4, 4,
			[][2]int{
				{1, 2},
				{1, 3},
				{2, 4},
				{3, 4},
			},
			3,
			[][2]int{
				{1, 2},
				{1, 3},
				{2, 3},
			},
			2,
		},

		{"入力例2",
			4, 4,
			[][2]int{
				{1, 2},
				{1, 3},
				{2, 4},
				{3, 4},
			},
			4,
			[][2]int{
				{3, 4},
				{1, 2},
				{2, 4},
				{2, 4},
			},
			4,
		},

		{"入力例3",
			6, 12,
			[][2]int{
				{2, 3},
				{4, 6},
				{1, 2},
				{4, 5},
				{2, 6},
				{1, 5},
				{4, 5},
				{1, 3},
				{1, 2},
				{2, 6},
				{2, 3},
				{2, 5},
			},
			5,
			[][2]int{
				{3, 5},
				{1, 4},
				{2, 6},
				{4, 6},
				{5, 6},
			},
			9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC190Cその2(tt.N, tt.M, tt.AB, tt.K, tt.CD)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
