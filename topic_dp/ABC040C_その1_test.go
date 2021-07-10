package topic_dp

import (
	"testing"
)

// [ABC040C - 柱柱柱柱柱](https://atcoder.jp/contests/abc040/tasks/abc040_c)
func AnswerABC040Cその1(N int, A []int) int {
	dp := make([]int, N)

	dp[0] = 0 // ゼロ値があるから不要だけど、手順になれるためにわざと書いておく

	for i := 1; i < N; i++ {
		if i == 1 {
			dp[i] = dp[0] + AbsInt(A[1]-A[0])
		} else {
			dp[i] = Min(
				dp[i-1]+AbsInt(A[i]-A[i-1]),
				dp[i-2]+AbsInt(A[i]-A[i-2]),
			)
		}
	}

	return dp[N-1]
}

func TestAnswerABC040Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want int
	}{
		{"入力例1", 4, []int{100, 150, 130, 120}, 40},
		{"入力例2", 4, []int{100, 125, 80, 110}, 40},
		{"入力例3", 9, []int{314, 159, 265, 358, 979, 323, 846, 264, 338}, 310},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC040Cその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
