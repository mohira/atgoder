package topic_dp

import (
	"testing"
)

// [EDPC B - Frog 2](https://atcoder.jp/contests/dp/tasks/dp_b)
func AnswerEDPCAFrog2(N, K int, H []int) int {
	return 0
}

func TestEDPCAFrog2(t *testing.T) {
	tests := []struct {
		name string
		N, K int
		H    []int
		want int
	}{
		{"入力例1", 5, 3, []int{10, 30, 40, 50, 20}, 30},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerEDPCAFrog2(tt.N, tt.K, tt.H)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
