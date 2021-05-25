package problemB

import (
	"math"
	"strconv"
	"testing"
)

// [ABC112B - Time Limit Exceeded](https://atcoder.jp/contests/abc112/tasks/abc112_b)
func AnswerABC112Bその1(N, T int, CT [][]int) string {
	var minCost = math.MaxInt64
	for _, ct := range CT {
		c, t := ct[0], ct[1]

		if t <= T {
			minCost = min(minCost, c)
		}
	}

	if minCost == math.MaxInt64 {
		return "TLE"
	} else {
		return strconv.Itoa(minCost)
	}
}

func TestAnswerABC112Bその1(t *testing.T) {
	tests := []struct {
		name string
		N, T int
		CT   [][]int
		want string
	}{
		{
			"入力例1",
			3, 70,
			[][]int{
				{7, 60},
				{1, 80},
				{4, 50},
			},
			"4",
		},
		{
			"入力例2",
			4, 3,
			[][]int{
				{1, 1000},
				{2, 4},
				{3, 1000},
				{4, 500},
			},
			"TLE",
		},
		{
			"入力例3",
			5, 9,
			[][]int{
				{25, 8},
				{5, 9},
				{4, 10},
				{1000, 1000},
				{6, 1},
			},
			"5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC112Bその1(tt.N, tt.T, tt.CT)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
