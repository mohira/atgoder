package similar

import (
	"sort"
	"testing"
)

// [ABC067B - Snake Toy](https://atcoder.jp/contests/abc067/tasks/abc067_b)
func AnswerABC067Bその1(N, K int, L []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(L)))

	var ans int

	for i := 0; i < K; i++ {
		ans += L[i]
	}

	return ans
}

func TestAnswerABC067Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		K    int
		L    []int
		want int
	}{
		{"入力例1", 5, 3, []int{1, 2, 3, 4, 5}, 12},
		{"入力例2", 15, 14, []int{50, 26, 27, 21, 41, 7, 42, 35, 7, 5, 5, 36, 39, 1, 45}, 386},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC067Bその1(tt.N, tt.K, tt.L)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}

		})

	}
}
