package ABC_problemB

import (
	"testing"
)

// [ABC174B - Distance](https://atcoder.jp/contests/abc174/tasks/abc174_b)
func AnswerABC174Bその1(N, D int, XY [][]int) int {
	var count int
	for _, xy := range XY {
		x, y := xy[0], xy[1]

		if x*x+y*y <= D*D {
			count++
		}
	}
	return count
}

func TestAnswerABC174Bその1(t *testing.T) {
	tests := []struct {
		name string
		N, D int
		XY   [][]int
		want int
	}{
		{"入力例1",
			4, 5,
			[][]int{
				{0, 5},
				{-2, 4},
				{3, 4},
				{4, -4},
			},
			3,
		},

		{"入力例2",
			12, 3,
			[][]int{
				{1, 1},
				{1, 1},
				{1, 1},
				{1, 1},
				{1, 2},
				{1, 3},
				{2, 1},
				{2, 2},
				{2, 3},
				{3, 1},
				{3, 2},
				{3, 3},
			},
			7,
		},

		{"入力例3",
			20, 100000,
			[][]int{
				{14309, -32939},
				{-56855, 100340},
				{151364, 25430},
				{103789, -113141},
				{147404, -136977},
				{-37006, -30929},
				{188810, -49557},
				{13419, 70401},
				{-88280, 165170},
				{-196399, 137941},
				{-176527, -61904},
				{46659, 115261},
				{-153551, 114185},
				{98784, -6820},
				{94111, -86268},
				{-30401, 61477},
				{-55056, 7872},
				{5901, -163796},
				{138819, -185986},
				{-69848, -96669},
			},
			6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC174Bその1(tt.N, tt.D, tt.XY)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
