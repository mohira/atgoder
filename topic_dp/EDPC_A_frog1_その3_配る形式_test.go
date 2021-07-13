package topic_dp

import (
	"math"
	"testing"
)

// [EDPC A - Frog 1](https://atcoder.jp/contests/dp/tasks/dp_a)
// 配る遷移形式
func AnswerEDPCAFrog1その3(N int, H []int) int {
	const INF = math.MaxInt64

	dp := make([]int, N)

	for i := 0; i < N; i++ {
		dp[i] = INF
	}

	dp[0] = 0

	for i := 0; i < N; i++ {
		// もらう形式では Indexのチェックが、 `0 <= i-1` だった
		if i+1 < N {
			ChMin(&dp[i+1], dp[i]+AbsInt(H[i]-H[i+1]))
		}
		if i+2 < N {
			ChMin(&dp[i+2], dp[i]+AbsInt(H[i]-H[i+2]))
		}
	}

	return dp[N-1]
}

func TestEDPCAFrog1その3(t *testing.T) {

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
			got := AnswerEDPCAFrog1その3(tt.N, tt.H)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}

		})
	}
}
