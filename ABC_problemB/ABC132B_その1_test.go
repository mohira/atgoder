package ABC_problemB

import (
	"testing"
)

// [ABC132B - Golden Apple](https://atcoder.jp/contests/abc132/tasks/abc132_b)
func AnswerABC132Bその1(n int, p []int) int {
	var count int

	for i := 0; i <= n-3; i++ {
		a, b, c := p[i], p[i+1], p[i+2]

		if a <= b && b <= c || c <= b && b <= a {
			count++
		}
	}

	return count
}

func TestAnswerABC132Bその1(t *testing.T) {
	tests := []struct {
		name string
		n    int
		p    []int
		want int
	}{
		{"入力例1", 5, []int{1, 3, 5, 4, 2}, 2},
		{"入力例2", 9, []int{9, 6, 3, 2, 5, 8, 7, 4, 1}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC132Bその1(tt.n, tt.p)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
