package ABC_problemC_gray

import (
	"atgoder/lib"
	"testing"
)

// [ABC170C - Forbidden List](https://atcoder.jp/contests/abc170/tasks/abc170_c)
func AnswerABC170Cその1(X, N int, P []int) int {
	// 解は[0, 101]なことに注意する
	bucket := make([]int, 102)
	for _, p := range P {
		bucket[p]++
	}
	bucket[X]++

	// Xが一票なら、X自身が答え
	if bucket[X] == 1 {
		return X
	}

	// X未満かつ、未登場な値を検索
	front := 0
	for i, v := range bucket[:X] {
		if v == 0 {
			front = i
		}
	}

	// Xより大きい かつ 未登場な値を検索
	back := 0
	for i, v := range bucket {
		if i > X && v == 0 {
			back = i
			break
		}
	}

	diff1 := lib.AbsInt(X - front)
	diff2 := lib.AbsInt(X - back)

	if diff1 < diff2 {
		return front
	} else if diff2 < diff1 {
		return back
	} else {
		return lib.Min(front, back)
	}

}

func TestAnswerABC170Cその1(t *testing.T) {
	tests := []struct {
		name string
		X, N int
		P    []int
		want int
	}{
		{"入力例1", 6, 5, []int{4, 7, 10, 6, 5}, 8},
		{"入力例2", 10, 5, []int{4, 7, 10, 6, 5}, 9},
		{"入力例3", 100, 0, []int{}, 100},
		{"答えが最大のケース", 100, 2, []int{99, 100}, 101},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC170Cその1(tt.X, tt.N, tt.P)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
