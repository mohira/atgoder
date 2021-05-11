package topic_cumlative_sum

import (
	"testing"
)

// [ABC037C - 総和](https://atcoder.jp/contests/abc037/tasks/abc037_c)
func AnswerABC037Cその1(N, K int, A []int) int {
	cumsum := make([]int, N+1)
	for i := 0; i < N; i++ {
		cumsum[i+1] = cumsum[i] + A[i]
	}

	var ans int
	for i := 0; i < len(cumsum)-K; i++ {
		ans += cumsum[i+K] - cumsum[i]
	}

	return ans
}

func TestAnswerABC037Cその1(t *testing.T) {
	tests := []struct {
		name string
		N, K int
		A    []int
		want int
	}{
		{"入力例1", 5, 3, []int{1, 2, 4, 8, 16}, 49},
		{"入力例2", 20, 10, []int{100000000, 100000000, 98667799, 100000000, 100000000, 100000000, 100000000, 99986657, 100000000, 100000000, 100000000, 100000000, 100000000, 98995577, 100000000, 100000000, 99999876, 100000000, 100000000, 99999999}, 10988865195},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC037Cその1(tt.N, tt.K, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
