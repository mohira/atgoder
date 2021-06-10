package ABC_problemB

import (
	"testing"
)

// [ABC125B - Resale](https://atcoder.jp/contests/abc125/tasks/abc125_b)
func AnswerABC125Bその1(N int, V, C []int) int {
	var ans int
	for i := 0; i < N; i++ {
		diff := V[i] - C[i]

		if 0 < diff {
			ans += diff
		}
	}
	return ans
}

func TestAnswerABC125Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		V    []int
		C    []int
		want int
	}{
		{
			"入力例1",
			3,
			[]int{10, 2, 5},
			[]int{6, 3, 4},
			5,
		},
		{
			"入力例2",
			4,
			[]int{13, 21, 6, 19},
			[]int{11, 30, 6, 15},
			6,
		},
		{
			"入力例3",
			1,
			[]int{1},
			[]int{50},
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC125Bその1(tt.N, tt.V, tt.C)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
