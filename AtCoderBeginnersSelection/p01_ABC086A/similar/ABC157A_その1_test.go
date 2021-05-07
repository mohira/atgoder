package similar

import (
	"testing"
)

// [ABC157A - Duplex Printing](https://atcoder.jp/contests/abc157/tasks/abc157_a)
func AnswerABC157Aその1(N int) int {
	return N/2 + N%2
}

func TestAnswerABC157Aその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{"入力例1", 5, 3},
		{"入力例2", 2, 1},
		{"入力例3", 100, 50},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC157Aその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
