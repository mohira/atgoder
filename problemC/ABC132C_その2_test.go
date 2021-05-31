package problemC

import (
	"atgoder/lib"
	"sort"
	"testing"
)

// [ABC132C - Divide the Problems](https://atcoder.jp/contests/abc132/tasks/abc132_c)
func AnswerABC132Cその2(N int, D []int) int {
	// わざわざ分割したスライスを使わなくても求められる
	// (分割したほうがほうがイメージはつかみやすい)
	sort.Ints(D)

	lower := D[N/2-1] + 1
	upper := D[N/2]

	return lib.Max(upper-lower+1, 0)
}

func TestAnswerABC132Cその2(t *testing.T) {
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
			got := AnswerABC132Cその2(tt.N, tt.D)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
