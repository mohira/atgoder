package ABC_problemC_gray

import (
	"atgoder/lib"
	"testing"
)

// [ABC176C - Step](https://atcoder.jp/contests/abc176/tasks/abc176_c)
func AnswerABC176Cその2(N int, A []int) int {
	var maxHeight int
	var ans int

	// 前から順番に最大の高さをキープしていく
	// 現状の最大の高さと当該の高さとの差分を調べていけばいい
	for i := 0; i < N; i++ {
		maxHeight = lib.Max(maxHeight, A[i])
		ans += maxHeight - A[i]
	}

	return ans
}

func TestAnswerABC176Cその2(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want int
	}{
		{"入力例1", 5, []int{2, 1, 5, 4, 3}, 4},
		{"入力例2", 5, []int{3, 3, 3, 3, 3}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC176Cその2(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
