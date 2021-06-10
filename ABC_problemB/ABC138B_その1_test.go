package ABC_problemB

import (
	"atgoder/lib"
	"math"
	"testing"
)

// [ABC138B - Power Socket](https://atcoder.jp/contests/abc138/tasks/abc138_b)
func AnswerABC138Bその1(N int, A []int) float64 {
	// ラストまでfloatを使わないパターン(最小公倍数を分母分子にかける)
	if N == 1 {
		return float64(A[0])
	}

	lcm := lib.LCM(A)

	denominator := 0
	for _, a := range A {
		denominator += lcm / a
	}

	return float64(lcm) / float64(denominator)
}

func TestAnswerABC138Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want float64
	}{
		{"入力例1", 2, []int{10, 30}, 7.5},
		{"入力例2", 3, []int{200, 200, 200}, 66.66666666666667},
		{"入力例3", 1, []int{1000}, 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC138Bその1(tt.N, tt.A)

			if math.Abs(got-tt.want) > 1e-5 {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
