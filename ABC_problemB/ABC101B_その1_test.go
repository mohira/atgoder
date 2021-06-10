package ABC_problemB

import (
	"testing"
)

// [ABC101B - Digit Sums](https://atcoder.jp/contests/abc101/tasks/abc101_b)
func AnswerABC101Bその1(N int) string {
	var S int

	tmp := N
	for tmp > 0 {
		S += tmp % 10
		tmp /= 10
	}

	if N%S == 0 {
		return "Yes"
	} else {
		return "No"
	}
}

func TestAnswerABC101Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want string
	}{
		{"入力例1", 12, "Yes"},
		{"入力例2", 101, "No"},
		{"入力例3", 999999999, "Yes"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC101Bその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
