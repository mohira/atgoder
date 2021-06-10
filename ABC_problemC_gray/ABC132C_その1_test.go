package ABC_problemC_gray

import (
	"atgoder/lib"
	"sort"
	"testing"
)

// [ABC132C - Divide the Problems](https://atcoder.jp/contests/abc132/tasks/abc132_c)
func AnswerABC132Cその1(N int, D []int) int {
	sort.Ints(D)

	for i := 1; i < N; i++ {
		left, right := D[:i], D[i:]

		if len(left) == len(right) {
			lower := left[len(left)-1] + 1
			upper := right[0]

			return lib.Max(upper-lower+1, 0)
		}
	}

	return 0
}

func TestAnswerABC132Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		D    []int
		want int
	}{
		{"入力例1", 6, []int{9, 1, 4, 4, 6, 7}, 2},
		{"入力例2", 8, []int{9, 1, 14, 5, 5, 4, 4, 14}, 0},
		{"入力例3", 14, []int{99592, 10342, 29105, 78532, 83018, 11639, 92015, 77204, 30914, 21912, 34519, 80835, 100000, 1}, 42685},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC132Cその1(tt.N, tt.D)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
