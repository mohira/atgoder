package ABC_problemC

import (
	"testing"
)

// [ABC177C - Sum of product of pairs](https://atcoder.jp/contests/abc177/tasks/abc177_c)
func AnswerABC177Cその1(N int, A []int) int {
	// 式変形をしていくと、Ai * 区間和 になっている
	// 区間は連続しているので累積和を使えばO(1)
	//
	// 上三角行列を見るイメージ
	// A1(A2+A3+A4+...+An)
	// A2(   A3+A4+...+An)
	// A3(      A4+...+An)
	const mod int = 1e+9 + 7

	cumsum := make([]int, N+1)
	for i := 0; i < N; i++ {
		cumsum[i+1] = (cumsum[i] + A[i]) % mod
	}

	var ans int
	for i := 0; i < N; i++ {
		rangeSum := (cumsum[N] - cumsum[i+1]) % mod

		// modを足すことで回避
		ans += A[i] * (rangeSum + mod) % mod
		ans %= mod
	}

	return ans
}

func TestAnswerABC177Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want int
	}{
		{"入力例1", 3, []int{1, 2, 3}, 11},
		{"入力例2", 4, []int{141421356, 17320508, 22360679, 244949}, 437235829},
		{"modをたさないとマイナスになるケース(意味わかってない)", len(data177C), data177C, 926527560},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC177Cその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
