package topic_dp

import (
	"testing"
)

// [ABC099C - Strange Bank](https://atcoder.jp/contests/abc099/tasks/abc099_c)
func AnswerABC099Cその1(N int) int {
	// dp[i] → i円の支払いに必要な最低の操作回数
	dp := make([]int, N+1)
	dp[0] = 0

	for i := 1; i <= N; i++ {
		dp[i] = 1 + dp[i-1] // 1円引き出し

		// 6^N 円引き出し
		for pow6 := 6; 0 <= i-pow6; pow6 *= 6 {
			dp[i] = Min(dp[i], 1+dp[i-pow6])
		}

		// 9^N 円引き出し
		for pow9 := 9; 0 <= i-pow9; pow9 *= 9 {
			dp[i] = Min(dp[i], 1+dp[i-pow9])
		}
	}

	return dp[N]
}

func TestAnswerABC099Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{"入力例1", 127, 4},
		{"入力例2", 3, 3},
		{"入力例3", 44852, 16},
		{"最大ケースで時間間に合うか？", 100000, 7},
		{"入力例1を小さくしたやつ", 46, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC099Cその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
