package ABC_problemB

import (
	"atgoder/lib"
	"math"
	"testing"
)

// [B - おいしいたこ焼きの食べ方](https://atcoder.jp/contests/abc005/tasks/abc005_2)
func AnswerABC005Bその1(N int, T []int) int {
	var min int = math.MaxInt64

	for _, t := range T {
		min = lib.Min(min, t)
	}

	return min
}

func TestAnswerABC005Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		T    []int
		want int
	}{
		{"入力例1", 3, []int{1, 2, 3}, 1},
		{"入力例2", 3, []int{3, 3, 3}, 3},
		{"入力例3", 5, []int{3, 1, 4, 1, 5}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC005Bその1(tt.N, tt.T)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
