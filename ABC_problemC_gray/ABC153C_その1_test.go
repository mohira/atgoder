package ABC_problemC_gray

import (
	"atgoder/lib"
	"sort"
	"testing"
)

// [ABC153C - Fennec vs Monster](https://atcoder.jp/contests/abc153/tasks/abc153_c)
func AnswerABC153Cその1(N, K int, H []int) int {
	// 降順に必殺技回数分の要素を落とせばいい
	sort.Ints(H)

	idx := lib.Max(N-K, 0)

	return lib.SumInts(H[:idx])
}

func TestAnswerABC153Cその1(t *testing.T) {
	tests := []struct {
		name string
		N, K int
		H    []int
		want int
	}{
		{
			"入力例1",
			3, 1,
			[]int{4, 1, 5},
			5,
		},
		{
			"入力例2",
			8, 9,
			[]int{7, 9, 3, 2, 3, 8, 4, 6},
			0,
		},
		{
			"入力例3",
			3, 0,
			[]int{1000000000, 1000000000, 1000000000},
			3000000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC153Cその1(tt.N, tt.K, tt.H)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
