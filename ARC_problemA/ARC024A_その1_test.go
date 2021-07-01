package ARC_problemA

import (
	"atgoder/lib"
	"testing"
)

// [ARC024A - くつがくっつく](https://atcoder.jp/contests/agc024/tasks/agc024_a)
func AnswerARC024Aその1(L, R int, ls, rs []int) int {
	bucket1 := make([]int, 41)
	for _, l := range ls {
		bucket1[l]++
	}
	bucket2 := make([]int, 41)
	for _, r := range rs {
		bucket2[r]++
	}

	var count int

	for i := 0; i <= 40; i++ {
		l := bucket1[i]
		r := bucket2[i]

		if 0 < l && 0 < r {
			count += lib.Min(l, r)
		}
	}

	return count
}

func TestAnswerARC024Aその1(t *testing.T) {
	tests := []struct {
		name   string
		L, R   int
		ls, rs []int
		want   int
	}{
		{
			"入力例1",
			3, 3,
			[]int{20, 21, 22},
			[]int{30, 22, 15},
			1,
		},
		{
			"入力例2",
			3, 4,
			[]int{10, 11, 10},
			[]int{12, 10, 11, 25},
			2,
		},
		{
			"入力例3",
			5, 5,
			[]int{10, 10, 10, 10, 10},
			[]int{10, 10, 10, 10, 10},
			5,
		},
		{
			"入力例4",
			5, 5,
			[]int{10, 11, 12, 13, 14},
			[]int{30, 31, 32, 33, 34},
			0,
		},

		{
			"Rがおおくて、あまるやつ",
			3, 4,
			[]int{10, 10, 10},
			[]int{10, 10, 10, 10},
			3,
		},
		{
			"Lがおおくて、あまるやつ",
			4, 3,
			[]int{10, 10, 10, 10},
			[]int{10, 10, 10},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerARC024Aその1(tt.L, tt.R, tt.ls, tt.rs)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
