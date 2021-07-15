package topic_dp

import (
	"testing"
)

// [EDPC E - Knapsack1](https://atcoder.jp/contests/dp/tasks/dp_e)
func AnswerEDPC問題Eその1(N, W int, WV [][]int) int {
	return 0
}

func TestAnswerEDPC問題Eその1(t *testing.T) {
	tests := []struct {
		name string
		N, W int
		WV   [][]int
		want int
	}{
		{
			"入力例1",
			3, 8,
			[][]int{
				{3, 30},
				{4, 50},
				{5, 60},
			},
			90,
		},
		{
			"入力例2",
			1, 1000000000,
			[][]int{
				{1000000000, 10},
			},
			10,
		},
		{
			"入力例3",
			6, 15,
			[][]int{
				{6, 5},
				{5, 6},
				{6, 4},
				{6, 6},
				{3, 5},
				{7, 2},
			},
			17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnswerEDPC問題Eその1(tt.N, tt.W, tt.WV); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
