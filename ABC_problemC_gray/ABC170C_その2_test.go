package ABC_problemC_gray

import (
	"atgoder/lib"
	"math"
	"testing"
)

// [ABC170C - Forbidden List](https://atcoder.jp/contests/abc170/tasks/abc170_c)
func AnswerABC170Cその2(X, N int, P []int) int {
	// [0, 101]で全探索しつつ、Xとの距離をみて更新する
	bucket := make([]int, 1000)
	for _, p := range P {
		bucket[p] = 1
	}

	var minDiff = math.MaxInt64
	var ans int

	for i := 0; i <= 101; i++ {
		notUsed := bucket[i] == 0

		diff := lib.AbsInt(i - X)

		if notUsed && (diff < minDiff) {
			minDiff = diff
			ans = i
		}
	}

	return ans
}

func TestAnswerABC170Cその2(t *testing.T) {
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
			got := AnswerABC170Cその2(tt.X, tt.N, tt.P)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
