package problemB

import (
	"math"
	"testing"
)

// [ABC129B - Balance](https://atcoder.jp/contests/abc130/tasks/abc130_b)
func AnswerABC129Bその1(N int, W []int) int {
	// 累積和を使えばO(N)
	cumsum := make([]int, N+1)
	for i := 1; i <= N; i++ {
		cumsum[i] = cumsum[i-1] + W[i-1]
	}

	total := sumInts(W)

	var minDiff = math.MaxInt64
	for i := 0; i < N; i++ {
		s1 := cumsum[i]
		s2 := total - cumsum[i]

		minDiff = min(minDiff, absInt(s1-s2))
	}

	return minDiff
}

func TestAnswerABC129Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		W    []int
		want int
	}{
		{"入力例1", 3, []int{1, 2, 3}, 0},
		{"入力例2", 4, []int{1, 3, 1, 1}, 2},
		{"入力例3", 8, []int{27, 23, 76, 2, 3, 5, 62, 52}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC129Bその1(tt.N, tt.W)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
