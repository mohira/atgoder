package ABC_problemC_gray

import (
	"testing"
)

// [ABC185C - Duodecim Ferra](https://atcoder.jp/contests/abc185/tasks/abc185_c)
func AnswerABC185Cその1(L int) int {
	// 長さ1間隔の点があるとして、そこから11個選べばいい => nC11
	// 点を打てる場所 は L-1 個 => L-1C11

	var ans = 1

	// 64bitでオーバーフローしないように、逐次割りながら計算を進める
	for i := 0; i < 11; i++ {
		ans *= (L - 1) - i
		ans /= i + 1
	}

	return ans
}

func TestAnswerABC185Cその1(t *testing.T) {
	tests := []struct {
		name string
		L    int
		want int
	}{
		{"入力例1", 12, 1},
		{"入力例2", 13, 12},
		{"入力例3", 17, 4368},
		{"MAX", 200, 366461620334848584},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC185Cその1(tt.L)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
