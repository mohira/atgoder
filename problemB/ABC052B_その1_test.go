package problemB

import (
	"testing"
)

// [ABC052B - Increment Decrement](https://atcoder.jp/contests/abc052/tasks/abc052_b)
func AnswerABC052Bその1(N int, S string) int {
	maxValue := 0
	x := 0
	for _, c := range S {
		if c == 'I' {
			x++
		} else {
			x--
		}

		maxValue = max(maxValue, x)
	}

	return maxValue
}

func TestAnswerABC052Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		S    string
		want int
	}{
		{"入力例1", 5, "IIDID", 2},
		{"入力例2", 7, "DDIDDII", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC052Bその1(tt.N, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
