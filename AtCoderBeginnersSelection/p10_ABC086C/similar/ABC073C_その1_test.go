package p10_ABC086C

import (
	"testing"
)

// [ABC073C - Write and Erase](https://atcoder.jp/contests/abc073/tasks/abc073_c)
func AnswerABC073Cその1(N int, A []int) int {
	var bucket = make(map[int]int)
	for _, a := range A {
		bucket[a]++
	}

	var count int

	for _, frequency := range bucket {
		if frequency%2 == 1 {
			count++
		}
	}

	return count
}

func TestAnswerABC073Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want int
	}{
		{"入力例1", 3, []int{6, 2, 6}, 1},
		{"入力例2", 4, []int{2, 5, 5, 2}, 0},
		{"入力例3", 6, []int{12, 22, 16, 22, 18, 12}, 2},

		// https://youtu.be/7ctPpofZdJ4?t=1696 からの解説が素晴らしい
		// [2,3,2,3,2] => 2が3個 3が2個
		// [2,2,2,3,2] => 2が3個 3が2個
		// 重要なのは個数情報！ もとの数列情報じゃないよ！
		{"", 5, []int{2, 3, 2, 3, 2}, 1},
		{"", 5, []int{2, 2, 2, 3, 3}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC073Cその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
