package ABC_problemB

import (
	"testing"
)

// [ABC056B - NarrowRectanglesEasy](https://atcoder.jp/contests/abc056/tasks/abc056_b)
func AnswerABC056Bその1(W, a, b int) int {
	// もっと簡約にできるだろうけど、重なりパターンを愚直に条件分岐
	if a < b {
		if a+W < b {
			return b - (a + W)
		} else {
			return 0
		}
	}

	if b < a {
		if b+W < a {
			return a - (b + W)
		} else {
			return 0
		}
	}

	return 0
}

func TestAnswerABC056Bその1(t *testing.T) {
	tests := []struct {
		name    string
		W, a, b int
		want    int
	}{
		{"入力例1: ___ ---", 3, 2, 6, 1},
		{"入力例2: ___ ---", 3, 1, 3, 0},
		{"入力例3: --- ___", 5, 10, 1, 4},
		{"入力例a: _==-", 3, 1, 3, 0},
		{"入力例b: -==_", 3, 4, 2, 0},
		{"入力例b: ===", 3, 3, 3, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC056Bその1(tt.W, tt.a, tt.b)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
