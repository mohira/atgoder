package topic_dp

import (
	"testing"
)

// [EDPC E - Knapsack1](https://atcoder.jp/contests/dp/tasks/dp_e)
func AnswerEDPC問題Eその1(N, W int, WV [][]int) int {
	const (
		maxN       int = 1e+2
		maxV       int = 1e+3
		maxAmountV     = maxN*maxV + 1 // 番兵あり
		maxW       int = 1e+9 + 1      // 番兵あり
		INF            = maxN * maxW
	)

	// dp[i][j]: 品物iまでで、価値がちょうどjとなるときの総重量の最小値
	// 初期条件
	// 	dp[i][0] =   0 // 価値がちょうど0なので、総重量も0
	// 	dp[0][j] = INF // 品物を1つも使わず、価値jは実現できない。なのでINF
	dp := make([][]int, N+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, maxAmountV)
	}

	for j := 1; j < maxAmountV; j++ {
		dp[0][j] = INF
	}

	for i := 1; i <= N; i++ {
		wi, vi := WV[i-1][0], WV[i-1][1]
		for j := 0; j < maxAmountV; j++ {
			dp[i][j] = dp[i-1][j] // 品物iを不採用

			if 0 <= j-vi {
				// 品物iを採用したら更新チャンス
				dp[i][j] = Min(dp[i][j], dp[i-1][j-vi]+wi)
			}
		}
	}

	var ans int
	// vを動かしていることに注意
	for v := 0; v < maxAmountV; v++ {
		w := dp[N][v] // 品物Nまでで、価値がvとなるときの「重さ」

		if w <= W {
			ans = Max(ans, v)
		}
	}
	return ans
}

func TestAnswerEDPC問題Eその1(t *testing.T) {
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
				{3, 3},
				{4, 5},
				{5, 6},
			},
			9,
		},
		{
			"入力例2",
			1, 1000000000,
			[][]int{
				{1000000000, 10},
			},
			10,
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
			if got := AnswerEDPC問題Eその1(tt.N, tt.W, tt.WV); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
