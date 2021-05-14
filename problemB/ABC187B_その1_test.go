package problemB

import (
	"math"
	"testing"
)

// [ABC187B - Gentle Pairs](https://atcoder.jp/contests/abc187/tasks/abc187_b)
func AnswerABC187Bその1(N int, XY [][]int) int {
	var count int

	for i := 0; i < N; i++ {
		xi, yi := XY[i][0], XY[i][1]
		for j := i + 1; j < N; j++ {
			xj, yj := XY[j][0], XY[j][1]

			katamuki := float64(yj-yi) / float64(xj-xi)
			if math.Abs(katamuki) <= 1 {
				count++
			}
		}
	}

	return count
}

func TestAnswerABC187Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		XY   [][]int
		want int
	}{
		{"入力例1",
			3,
			[][]int{
				{0, 0},
				{1, 2},
				{2, 1},
			},
			2,
		},
		{"入力例2",
			1,
			[][]int{
				{-691, 273},
			},
			0,
		},
		{"入力例3",
			10,
			[][]int{
				{-31, -35},
				{8, -36},
				{22, 64},
				{5, 73},
				{-14, 8},
				{18, -58},
				{-41, -85},
				{1, -88},
				{-21, -85},
				{-11, 82},
			},
			11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC187Bその1(tt.N, tt.XY)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
