package ABC_problemB

import (
	"testing"
)

// [ABC110B - 1 Dimensional World's Tale](https://atcoder.jp/contests/abc110/tasks/abc110_b)
func AnswerABC110Bその1(N, M, X, Y int, xs, ys []int) string {
	for Z := X + 1; Z <= Y; Z++ {
		ok := true
		for _, x := range xs {
			if !(x < Z) {
				ok = false
			}
		}
		for _, y := range ys {
			if !(y >= Z) {
				ok = false
			}
		}

		if ok {
			return "No War"
		}
	}

	return "War"
}

func TestAnswerABC110Bその1(t *testing.T) {
	tests := []struct {
		name       string
		N, M, X, Y int
		xs         []int
		ys         []int
		want       string
	}{
		{
			"入力例1",
			3, 2, 10, 20,
			[]int{8, 15, 13},
			[]int{16, 22},
			"No War",
		},
		{
			"入力例2",
			4, 2, -48, -1,
			[]int{-20, -35, -91, -23},
			[]int{-22, 66},
			"War",
		},
		{
			"入力例3",
			5, 3, 6, 8,
			[]int{-10, 3, 1, 5, -100},
			[]int{100, 6, 14},
			"War",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC110Bその1(tt.N, tt.M, tt.X, tt.Y, tt.xs, tt.ys)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
