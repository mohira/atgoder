package problemC

import (
	"atgoder/lib"
	"testing"
)

// [ABC139C - Lower](https://atcoder.jp/contests/abc139/tasks/abc139_c)
func AnswerABC139Cその1(N int, H []int) int {
	// 左右の高さの差分をみて、マイナス以下が連続で何度続くかを調べる
	var diffs = make([]int, N-1)
	for i := 0; i < N-1; i++ {
		diffs[i] = H[i+1] - H[i]
	}

	var steps int
	var ans int

	for _, d := range diffs {
		if d <= 0 {
			steps++
			ans = lib.Max(ans, steps)
		} else {
			ans = lib.Max(ans, steps)
			steps = 0
		}
	}

	return ans
}

func TestAnswerABC139Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		H    []int
		want int
	}{
		{"入力例1", 5, []int{10, 4, 8, 7, 3}, 2},
		{"入力例2", 7, []int{4, 4, 5, 6, 6, 5, 5}, 3},
		{"入力例3", 4, []int{1, 2, 3, 4}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC139Cその1(tt.N, tt.H)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
