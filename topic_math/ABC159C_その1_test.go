package topic_math

import (
	"math"
	"testing"
)

// [ABC159C - Maximum Volume](https://atcoder.jp/contests/abc159/tasks/abc159_c)
func AnswerABC159Cその1(L int) float64 {
	return math.Pow(float64(L)/3, 3)
}

func TestAnswerABC159Cその1(t *testing.T) {
	tests := []struct {
		name string
		L    int
		want float64
	}{
		{"入力例1", 3, 1.000000000000},
		{"入力例2", 999, 36926037.000000000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC159Cその1(tt.L)

			if math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
