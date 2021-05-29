package problemB

import (
	"atgoder/lib"
	"testing"
)

// [ABC074B - Collecting Balls (Easy Version)](https://atcoder.jp/contests/abc077/tasks/abc077_b)
func AnswerABC074Bその1(N, K int, X []int) int {
	var ans int

	for _, x := range X {
		da := 2 * lib.AbsInt(x-0)
		db := 2 * lib.AbsInt(x-K)
		ans += lib.Min(da, db)
	}

	return ans
}

func TestAnswerABC074Bその1(t *testing.T) {
	tests := []struct {
		name string
		N, K int
		X    []int
		want int
	}{
		{"入力例1", 1, 10, []int{2}, 4},
		{"入力例2", 2, 9, []int{3, 6}, 12},
		{"入力例3", 5, 20, []int{11, 12, 9, 17, 12}, 74},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC074Bその1(tt.N, tt.K, tt.X)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
