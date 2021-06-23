package ARC_problemA

import (
	"testing"
)

// [ARC006A - 宝くじ](https://atcoder.jp/contests/arc006/tasks/arc006_1)
func AnswerARC006Aその1(E []int, B int, L []int) int {
	var matchCount int
	for _, l := range L {
		for _, e := range E {
			if l == e {
				matchCount++
			}
		}
	}

	var containsBonusNumber bool
	for _, l := range L {
		if B == l {
			containsBonusNumber = true
		}
	}

	var ans int

	switch {
	case matchCount == 6:
		ans = 1
	case matchCount == 5 && containsBonusNumber:
		ans = 2
	case matchCount == 5:
		ans = 3
	case matchCount == 4:
		ans = 4
	case matchCount == 3:
		ans = 5
	}

	return ans
}

func TestAnswerARC006Aその1(t *testing.T) {
	tests := []struct {
		name string
		E    []int
		B    int
		L    []int
		want int
	}{
		{
			"入力例1",
			[]int{1, 2, 3, 4, 5, 6},
			7,
			[]int{1, 2, 3, 4, 5, 6},
			1,
		},
		{
			"入力例2",
			[]int{0, 1, 3, 5, 7, 9},
			4,
			[]int{0, 2, 4, 6, 8, 9},
			0,
		},
		{
			"入力例3",
			[]int{0, 2, 6, 7, 8, 9},
			4,
			[]int{0, 5, 6, 7, 8, 9},
			3,
		},
		{
			"入力例4",
			[]int{1, 3, 5, 6, 7, 8},
			9,
			[]int{3, 5, 6, 7, 8, 9},
			2,
		},
		{
			"入力例5",
			[]int{0, 1, 3, 4, 5, 7},
			8,
			[]int{2, 3, 5, 7, 8, 9},
			5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerARC006Aその1(tt.E, tt.B, tt.L)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
