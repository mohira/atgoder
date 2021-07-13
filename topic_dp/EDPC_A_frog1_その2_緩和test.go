package topic_dp

import (
	"math"
	"testing"
)

// [EDPC A - Frog 1](https://atcoder.jp/contests/dp/tasks/dp_a)
func AnswerEDPCAFrog1その2(N int, H []int) int {
	const INF = math.MaxInt64

	dp := make([]int, N)

	for i := 0; i < N; i++ {
		dp[i] = INF
	}

	dp[0] = 0

	for i := 1; i < N; i++ {
		ChMin(&dp[i], dp[i-1]+AbsInt(H[i]-H[i-1]))

		if i > 1 {
			ChMin(&dp[i], dp[i-2]+AbsInt(H[i]-H[i-2]))
		}
	}

	return dp[N-1]
}

func TestEDPCAFrog1その2(t *testing.T) {

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
			got := AnswerEDPCAFrog1その2(tt.N, tt.H)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}

		})
	}
}
