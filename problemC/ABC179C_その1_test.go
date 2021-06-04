package problemC

import (
	"testing"
)

// [ABC179C - A x B + C](https://atcoder.jp/contests/abc179/tasks/abc179_c)
func AnswerABC179Cその1(N int) int {
	// Aを固定するとBも決まるので、実はBの探索は不要
	var count int
	for A := 1; A < N; A++ {
		for B := 1; A*B < N; B++ {
			C := N - A*B
			if 1 <= C {
				count++
			}
		}
	}
	return count
}

func TestAnswerABC179Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{"入力例1", 3, 3},
		{"入力例2", 100, 473},
		{"入力例3", 1000000, 13969985},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC179Cその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
