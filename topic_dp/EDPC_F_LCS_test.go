package topic_dp

import (
	"testing"
)

// [EDPC F - LCS](https://atcoder.jp/contests/dp/tasks/dp_f)
func AnswerEDPC問題Fその1(s, t string) string {
	// dp[i][j]
	//	 	sのi文字目までと
	// 		tのj文字目までの
	// 	LCSの長さ
	dp := make([][]int, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(t)+1)
	}

	n := len(s)
	m := len(t)

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = Max(
				dp[i-1][j], // 左マスの値
				dp[i][j-1], // 上マスの値
			)

			if s[i-1] == t[j-1] {
				dp[i][j] = Max(
					dp[i][j],
					dp[i-1][j-1]+1, // 左上マス + 1
				)
			}
		}
	}

	var rewsna string
	for n > 0 && m > 0 {
		// dp[n][m]は表の右下(==すべてチェックしたあと)
		// 表を逆順にたどっていく
		// 逆順にたどるにはLCSの値をチェックすればよい
		// 	左マスから来た: dp[n][m] == dp[n-1][m]:
		// 	上マスから来た: dp[n][m] == dp[n][m-1]:
		// 	左上マスから来た: dp[n][m] == dp[n-1][m-1]+1:
		// で、左上マスから来た場合は、「その文字を選んでいる」ことになる
		// 「その文字」は s[n-1] または t[m-1] ← どっちでもOK

		switch {
		case dp[n][m] == dp[n-1][m]:
			// 1マス上をみて、そのときとLCSが同じ
			// => 上から来たことがわかる
			n--

		case dp[n][m] == dp[n][m-1]:
			// 1マス左をみて、そのときとLCSが同じ
			// => 左から来たことがわかる
			m--

		case dp[n][m] == dp[n-1][m-1]+1:
			// 左上をみて、そのときとLCSが増えた(+1)場合
			// => 左上から来たことがわかる
			rewsna += string(s[n-1]) // t[m-1]でも当然OK
			n--
			m--
		}
	}

	var ans string
	for i := len(rewsna) - 1; i >= 0; i-- {
		ans += string(rewsna[i])
	}
	return ans
}

func TestAnswerEDPC問題Fその1(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{
			"入力例1",
			"axyb",
			"abyxb",
			"axb",
		},
		{
			"入力例2",
			"aa",
			"xayaz",
			"aa",
		},
		{
			"入力例3",
			"a",
			"z",
			"",
		},
		{
			"入力例4",
			"abracadabra",
			"avadakedavra",
			"aaadara",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnswerEDPC問題Fその1(tt.s, tt.t); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
