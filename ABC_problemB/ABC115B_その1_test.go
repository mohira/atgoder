package ABC_problemB

import (
	"atgoder/lib"
	"testing"
)

// [ABC115B - Christmas Eve Eve](https://atcoder.jp/contests/abc115/tasks/abc115_b)
func AnswerABC115Bその1(N int, P []int) int {
	return lib.SumInts(P) - lib.FindMax(P)/2
}

func TestAnswerABC115Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		P    []int
		want int
	}{
		{"入力例1", 4, []int{4980, 7980, 6980}, 15950},
		{"入力例1", 4, []int{4320, 4320, 4320, 4320}, 15120},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC115Bその1(tt.N, tt.P)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
