package problemB

import (
	"testing"
)

// [ABC095B - Toll Gates](https://atcoder.jp/contests/abc095/tasks/abc095_b)
func AnswerABC095Bその1(N, M, X int, A []int) int {
	leftGates := 0
	rightGates := 0

	for _, a := range A {
		if a < X {
			leftGates++
		} else {
			rightGates++
		}
	}

	return min(leftGates, rightGates)
}

func TestAnswerABC095Bその1(t *testing.T) {
	tests := []struct {
		name    string
		N, M, X int
		A       []int
		want    int
	}{
		{
			"入力例1",
			5, 3, 3,
			[]int{1, 2, 4},
			1,
		},
		{
			"入力例2",
			7, 3, 2,
			[]int{4, 5, 6},
			0,
		},
		{
			"入力例3",
			10, 7, 5,
			[]int{1, 2, 3, 4, 6, 8, 9},
			3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC095Bその1(tt.N, tt.M, tt.X, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
