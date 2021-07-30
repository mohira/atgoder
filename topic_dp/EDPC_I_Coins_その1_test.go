package topic_dp

import (
	"math"
	"testing"
)

// [EDPC I - Coins](https://atcoder.jp/contests/dp/tasks/dp_i)
func AnswerEDPC問題Iその1(N int, P []float64) float64 {
	// dp[i][j] = コインを i枚 投げて、「表」が j回 でる確率
	dp := make([][]float64, N+1) // 0<=i<=N だよ
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]float64, N+1) //
	}

	dp[0][0] = 1 // 0枚投げて、「表」が0回となる確率は1

	for i := 1; i <= N; i++ {
		p := P[i-1]               // 1-indexなので注意
		for j := 0; j <= i; j++ { // 「表」は最大でも投げた回数しか出現しないので、 0<=j<=i
			// 「裏」の遷移は絶対起きる(
			dp[i][j] += dp[i-1][j] * (1 - p)

			// 「表」の遷移は起きない場合がある
			if j-1 >= 0 {
				dp[i][j] += dp[i-1][j-1] * p
			}
		}
	}
	// ex1
	// [1     0     0     0    ]
	// [0.7   0.3   0     0    ]
	// [0.28  0.54  0.18  0    ]
	// [0.056 0.332 0.468 0.144]
	//              |||||||||||| 「表」が「裏」より多く出ているところ

	var ans float64

	// dp[N] ⇔ 最終行 ⇔ i=N ⇔ 全てのコインを投げきった状態 だけを考えればよい
	for _, p := range dp[N][(N/2)+1:] {
		ans += p
	}

	return ans
}

func TestAnswerEDPC問題Iその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		P    []float64
		want float64
	}{
		{"入力例1", 3, []float64{0.30, 0.60, 0.80}, 0.612},
		{"入力例2", 1, []float64{0.50}, 0.5},
		{"入力例3", 5, []float64{0.42, 0.01, 0.42, 0.99, 0.42}, 0.3821815872},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerEDPC問題Iその1(tt.N, tt.P)
			if math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
