package topic_dp

import (
	"testing"
)

// [ABC129C - Typical Stairs](https://atcoder.jp/contests/abc129/tasks/abc129_c)
func AnswerABC129Cその3(N, M int, A []int) int {
	// i段目ではなく、(i-1)段目と(i-2)段目の床を気にする場合の実装
	const MOD int = 1e+9 + 7
	dp := make([]int, N+1)

	// i段目の床が使えるかどうか情報
	oks := make([]bool, N+1)
	for i := 0; i < N+1; i++ {
		oks[i] = true
	}
	for _, a := range A {
		oks[a] = false
	}

	// 初期条件
	dp[0] = 1
	if oks[1] {
		dp[1] = 1 // else条件はゼロ値があるので不要
	}

	// 前の床を気にする場合の実装(漸化式が見えにくいと思う)
	for i := 2; i <= N; i++ {
		if oks[i-1] {
			dp[i] += dp[i-1]
		}

		if oks[i-2] {
			dp[i] += dp[i-2]
		}

		dp[i] %= MOD
	}

	return dp[N]
}

func TestAnswerABC129Cその3(t *testing.T) {
	tests := []struct {
		name string
		N, M int
		A    []int
		want int
	}{
		{
			"入力例1",
			6, 1,
			[]int{3},
			4,
		},
		{
			"入力例1",
			10, 2,
			[]int{4, 5},
			0,
		},
		{
			"入力例1",
			100, 5,
			[]int{1, 23, 45, 67, 89},
			608200469,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC129Cその3(tt.N, tt.M, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
