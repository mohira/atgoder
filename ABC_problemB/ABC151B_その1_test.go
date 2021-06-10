package ABC_problemB

import (
	"atgoder/lib"
	"testing"
)

// [ABC151B - Achieve the Goal](https://atcoder.jp/contests/abc151/tasks/abc151_b)
func AnswerABC151Bその1(N, K, M int, A []int) int {
	var total int
	for _, a := range A {
		total += a
	}

	diff := N*M - total

	if diff <= K {
		return lib.Max(diff, 0)
	} else {
		return -1
	}
}

func TestAnswerABC151Bその1(t *testing.T) {
	tests := []struct {
		name    string
		N, K, M int
		A       []int
		want    int
	}{
		{"入力例1", 5, 10, 7, []int{8, 10, 3, 6}, 8},
		{"入力例2", 4, 100, 60, []int{100, 100, 100}, 0},
		{"入力例3", 4, 100, 60, []int{0, 0, 0}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC151Bその1(tt.N, tt.K, tt.M, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
