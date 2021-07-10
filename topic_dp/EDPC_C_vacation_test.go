package topic_dp

import (
	"testing"
)

// [EDPC C - Vacation](https://atcoder.jp/contests/dp/tasks/dp_c)
func AnswerEDPCCその1(N int, ABC [][]int) int {
	return 0
}

func TestAnswerEDPCCその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		ABC  [][]int
		want int
	}{
		{
			"入力例1",
			3,
			[][]int{
				{10, 40, 70},
				{20, 50, 80},
				{30, 60, 90},
			},
			210,
		},
		{
			"入力例2",
			1,
			[][]int{
				{100, 10, 1},
			},
			100,
		},
		{
			"入力例1",
			7,
			[][]int{
				{6, 7, 8},
				{8, 8, 3},
				{2, 5, 2},
				{7, 8, 6},
				{4, 6, 8},
				{2, 3, 4},
				{7, 5, 1},
			},
			46,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnswerEDPCCその1(tt.N, tt.ABC); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
