package ABC_problemB

import (
	"testing"
)

// [ABC179B - Go to Jail](https://atcoder.jp/contests/abc179/tasks/abc179_b)
func AnswerABC179Bその1(N int, D [][]int) string {
	var zoroCount int
	for i := 0; i < N; i++ {
		d1, d2 := D[i][0], D[i][1]

		if d1 == d2 {
			zoroCount++

			if zoroCount >= 3 {
				return "Yes"
			}
		} else {
			zoroCount = 0
		}
	}

	return "No"
}

func TestAnswerABC179Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		D    [][]int
		want string
	}{
		{"入力例1",
			5,
			[][]int{
				{1, 2},
				{6, 6},
				{4, 4},
				{3, 3},
				{3, 2},
			},
			"Yes",
		},
		{"入力例2",
			5,
			[][]int{
				{1, 1},
				{2, 2},
				{3, 4},
				{5, 5},
				{6, 6},
			},
			"No",
		},
		{"入力例3",
			6,
			[][]int{
				{1, 1},
				{2, 2},
				{3, 3},
				{5, 5},
				{6, 6},
			},
			"Yes",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC179Bその1(tt.N, tt.D)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
