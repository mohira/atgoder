package ARC_problemA

import (
	"atgoder/lib"
	"testing"
)

// [ARC036A - ぐっすり](https://atcoder.jp/contests/arc036/tasks/arc036_a)
func AnswerARC036Aその1(N int, K int, T []int) int {
	for i := 0; i <= N-3; i++ {
		sleepTime3Days := lib.SumInts(T[i : i+3])

		if sleepTime3Days < K {
			return i + 3
		}
	}
	return -1
}

func TestAnswerARC036Aその1(t *testing.T) {
	tests := []struct {
		name string
		N, K int
		T    []int
		want int
	}{
		{
			"入力例1",
			5, 1080,
			[]int{300, 420, 420, 180, 360},
			4,
		},
		{
			"入力例2",
			5, 180,
			[]int{60, 60, 60, 60, 60},
			-1,
		},
		{
			"入力例3",
			5, 4230,
			[]int{360, 360, 360, 360, 360},
			3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerARC036Aその1(tt.N, tt.K, tt.T)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
