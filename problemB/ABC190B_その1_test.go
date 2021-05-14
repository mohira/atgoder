package problemB

import (
	"testing"
)

// [ABC190B - Magic 3](https://atcoder.jp/contests/abc190/tasks/abc190_b)
func AnswerABC190Bその1(N, S, D int, XY [][]int) string {
	for i := 0; i < N; i++ {
		x, y := XY[i][0], XY[i][1]

		if x < S && D < y {
			return "Yes"
		}
	}
	return "No"
}

func TestAnswerABC190Bその1(t *testing.T) {
	tests := []struct {
		name    string
		N, S, D int
		XY      [][]int
		want    string
	}{
		{"入力例1",
			4, 9, 9,
			[][]int{
				{5, 5},
				{15, 5},
				{5, 15},
				{15, 15},
			},
			"Yes"},
		{"入力例2",
			3, 691, 273,
			[][]int{
				{691, 997},
				{593, 273},
				{691, 273},
			},
			"No"},
		{"入力例3",
			7, 100, 100,
			[][]int{
				{10, 11},
				{12, 67},
				{192, 79},
				{154, 197},
				{142, 158},
				{20, 25},
				{17, 108},
			},
			"Yes"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC190Bその1(tt.N, tt.S, tt.D, tt.XY)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
