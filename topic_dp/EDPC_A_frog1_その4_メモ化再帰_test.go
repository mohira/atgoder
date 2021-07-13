package topic_dp

import (
	"math"
	"testing"
)

// [EDPC A - Frog 1](https://atcoder.jp/contests/dp/tasks/dp_a)
// メモ化再帰
const INF = math.MaxInt64

func AnswerEDPCAFrog1その4(N int, H []int) int {
	dp := make([]int, N)
	for i := 0; i < N; i++ {
		dp[i] = INF
	}

	return rec(N-1, dp, H)
}

func rec(i int, dp []int, H []int) int {
	res := INF
	// DPの値が更新されていたらそのままリターン
	if dp[i] < INF {
		return dp[i]
	}

	// ベースケース: 足場0のコストは0
	if i == 0 {
		return 0
	}

	// 足場 i-1 から来た場合
	if 0 <= i-1 {
		cost1 := rec(i-1, dp, H) + AbsInt(H[i]-H[i-1])
		ChMin(&res, cost1)
	}

	// 足場 i-2 から来た場合
	if 0 <= i-2 {
		cost2 := rec(i-2, dp, H) + AbsInt(H[i]-H[i-2])
		ChMin(&res, cost2)
	}

	// 結果をメモする
	dp[i] = res

	return res
}
func TestEDPCAFrog1その4(t *testing.T) {

	tests := []struct {
		name string
		N    int
		H    []int
		want int
	}{
		{"入力例1", 4, []int{10, 30, 40, 20}, 30},
		{"入力例2", 2, []int{10, 10}, 0},
		{"入力例3", 6, []int{30, 10, 60, 10, 60, 50}, 40},
		{"けんちょん本", 7, []int{2, 9, 4, 5, 1, 6, 10}, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerEDPCAFrog1その4(tt.N, tt.H)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}

		})
	}
}
