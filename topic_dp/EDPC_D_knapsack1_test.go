package topic_dp

import (
	"testing"
)

// [EDPC D - Knapsack1](https://atcoder.jp/contests/dp/tasks/dp_d)
func AnswerEDPC問題Dその1(N, W int, WV [][]int) int {
	var dp [110][100010]int

	for i := 1; i <= N; i++ {
		wi := WV[i-1][0]
		vi := WV[i-1][1]

		for j := 0; j <= W; j++ {
			dp[i][j] = dp[i-1][j] // 品物i不採用なのでそのまま

			if 0 <= j-wi {
				// 品物iを採用する場合は、不採用の場合と比べて価値が大きい方
				dp[i][j] = Max(
					dp[i][j],
					dp[i-1][j-wi]+vi,
				)
			}
		}
	}

	return dp[N][W]
}

func TestAnswerEDPC問題Dその1(t *testing.T) {
	tests := []struct {
		name string
		N, W int
		WV   [][]int
		want int
	}{
		{
			"入力例1",
			3, 8,
			[][]int{
				{3, 30},
				{4, 50},
				{5, 60},
			},
			90,
		},
		{
			"入力例2",
			5, 5,
			[][]int{
				{1, 1000000000},
				{1, 1000000000},
				{1, 1000000000},
				{1, 1000000000},
				{1, 1000000000},
			},
			5000000000,
		},
		{
			"入力例3",
			6, 15,
			[][]int{
				{6, 5},
				{5, 6},
				{6, 4},
				{6, 6},
				{3, 5},
				{7, 2},
			},
			17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnswerEDPC問題Dその1(tt.N, tt.W, tt.WV); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
