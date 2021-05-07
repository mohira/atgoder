package topic_counter

import (
	"testing"
)

// [ABC118B - Foods Loved by Everyone](https://atcoder.jp/contests/abc118/tasks/abc118_b)
func AnswerABC118Bその1(N, M int, KA [][]int) int {
	bucket := make(map[int]int)
	for _, ka := range KA {
		A := ka[1:]
		for _, a := range A {
			bucket[a]++
		}
	}

	var count int

	for _, v := range bucket {
		// 人数分の投票があれば、みんな大好き
		if v == N {
			count++
		}
	}

	return count
}

func TestAnswerABC118Bその1(t *testing.T) {
	tests := []struct {
		name string
		N, M int
		KA   [][]int
		want int
	}{
		{"入力例1", 3, 4,
			[][]int{
				{2, 1, 3},
				{3, 1, 2, 3},
				{2, 3, 2},
			},
			1,
		},
		{"入力例2", 5, 5,
			[][]int{
				{4, 2, 3, 4, 5},
				{4, 1, 3, 4, 5},
				{4, 1, 2, 4, 5},
				{4, 1, 2, 3, 5},
				{4, 1, 2, 3, 4},
			},
			0,
		},
		{"入力例3", 1, 30,
			[][]int{
				{3, 5, 10, 30},
			},
			3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC118Bその1(tt.N, tt.M, tt.KA)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
