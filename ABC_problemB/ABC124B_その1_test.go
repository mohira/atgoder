package ABC_problemB

import (
	"atgoder/lib"
	"testing"
)

// [ABC124B - Great Ocean View](https://atcoder.jp/contests/abc124/tasks/abc124_b)
func AnswerABC124Bその1(N int, H []int) int {
	// 最大の高さをキープしていけばO(N)
	var count int
	var maxH int

	for i := 0; i < N; i++ {
		h := H[i]

		if maxH <= h {
			count++
		}

		maxH = lib.Max(maxH, h)
	}

	return count
}

func TestAnswerABC124Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		H    []int
		want int
	}{
		{"入力例1", 4, []int{6, 5, 6, 8}, 3},
		{"入力例2", 5, []int{4, 5, 3, 5, 4}, 3},
		{"入力例3", 5, []int{9, 5, 6, 8, 4}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC124Bその1(tt.N, tt.H)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
