package ABC_problemB

import (
	"testing"
)

// [ABC142B - Roller Coaster](https://atcoder.jp/contests/abc142/tasks/abc142_b)
func AnswerABC142Bその1(N, K int, H []int) int {
	var count int

	for _, h := range H {
		if K <= h {
			count++
		}
	}

	return count
}

func TestAnswerABC142Bその1(t *testing.T) {
	tests := []struct {
		name string
		N, K int
		D    []int
		want int
	}{
		{"入力例1", 4, 150, []int{150, 140, 100, 200}, 2},
		{"入力例2", 1, 500, []int{499}, 0},
		{"入力例3", 5, 1, []int{100, 200, 300, 400}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC142Bその1(tt.N, tt.K, tt.D)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
