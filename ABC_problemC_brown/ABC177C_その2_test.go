package ABC_problemC_brown

import (
	"testing"
)

// [ABC177C - Sum of product of pairs](https://atcoder.jp/contests/abc177/tasks/abc177_c)
func AnswerABC177Cその2(N int, A []int) int {
	// 累積和テーブルを使わず保存していくやり方
	//
	// 下三角行列を見るイメージ
	// A1(               )
	// A2(A1             ) ← カッコの中は次にも使える
	// A3(A1+A2          ) ← カッコの中は次にも使える
	// A4(A1+A2+A3       ) ← カッコの中は次にも使える
	// A5(A1+A2+A3+A4    ) ← カッコの中は次にも使える
	const mod int = 1e+9 + 7

	var ans int

	x := 0
	for i := 0; i < N; i++ {
		// modのタイミング注意
		// ans += (A[i] * x) % mod だけで、 ans%=mod しないとだめ
		// 1行で書くなら ans = (ans + A[i]*x) % mod
		ans += (A[i] * x) % mod
		ans %= mod

		x += A[i] % mod
		x %= mod
	}

	return ans
}

func TestAnswerABC177Cその2(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want int
	}{
		{"解説の例", 3, []int{3, 1, 4}, 19},
		{"入力例1", 3, []int{1, 2, 3}, 11},
		{"入力例2", 4, []int{141421356, 17320508, 22360679, 244949}, 437235829},
		{"modをたさないとマイナスになるケース(意味わかってない)", len(data177C), data177C, 926527560},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC177Cその2(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
