package similar

import (
	"math"
	"testing"
)

// [ABC113B - Palace](https://atcoder.jp/contests/abc113/tasks/abc113_b)
func AnswerABC113Bその1(T int, A int, H []int) int {
	var ansIdx int
	var minDiff = 1_000_000_000.0

	for i, h := range H {
		avgTemperature := float64(T) - float64(h)*0.006

		diff := math.Abs(float64(A) - avgTemperature)

		if diff < minDiff {
			ansIdx = i + 1
			minDiff = diff
		}
	}

	return ansIdx
}

func TestAnswerABC113Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		T    int
		A    int
		H    []int
		want int
	}{
		{"入力例1", 2, 12, 5, []int{1000, 2000}, 1},
		{"入力例2", 3, 21, -11, []int{81234, 94124, 52141}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC113Bその1(tt.T, tt.A, tt.H)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}

		})

	}
}
