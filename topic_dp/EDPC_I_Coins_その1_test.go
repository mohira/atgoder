package topic_dp

import (
	"math"
	"testing"
)

// [EDPC I - Coins](https://atcoder.jp/contests/dp/tasks/dp_i)
func AnswerEDPC問題Iその1(N int, P []float64) float64 {
	return 0
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
		{"入力例3", 1, []float64{0.42, 0.01, 0.42, 0.99, 0.42}, 0.3821815872},
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
