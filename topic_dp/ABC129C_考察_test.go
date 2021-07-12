package topic_dp

import (
	"testing"
)

// [ABC129C - Typical Stairs](https://atcoder.jp/contests/abc129/tasks/abc129_c)
// 考察として、「床が壊れていない場合」を考える
// [AtCoder ABC 129 C - Typical Stairs (茶色, 300 点) - けんちょんの競プロ精進記録](https://drken1215.hatenablog.com/entry/2019/06/10/140000)
func AnswerABC129C考察(N int) int {
	dp := make([]int, N+1)

	// 初期条件
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= N; i++ {
		// 「1段前から1マス移動する」 + 「2段前から2マス移動する」
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[N]
}
func TestAnswerABC129C考察(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{"", 1, 1},
		{"", 2, 2},
		{"", 3, 3},
		{"", 4, 5},
		{"", 5, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC129C考察(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
