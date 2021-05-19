package problemB

import (
	"testing"
)

// [ABC166B -  Trick or Treat](https://atcoder.jp/contests/abc166/tasks/abc166_b)
func AnswerABC166Bその1(N, K int, A [][]int) int {
	hasSnacks := make([]bool, N)

	for _, sunukeList := range A {
		for _, sunekeId := range sunukeList {
			sunekeId--
			hasSnacks[sunekeId] = true
		}
	}

	var count int
	for _, hasSnack := range hasSnacks {
		if !hasSnack {
			count++
		}
	}

	return count
}

func TestAnswerABC166Bその1(t *testing.T) {
	tests := []struct {
		name string
		N, K int
		A    [][]int
		want int
	}{
		{
			"入力例1",
			3, 2,
			[][]int{
				{1, 3},
				{3},
			},
			1,
		},
		{
			"入力例2",
			3, 3,
			[][]int{
				{3},
				{3},
				{3},
			},
			2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC166Bその1(tt.N, tt.K, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
