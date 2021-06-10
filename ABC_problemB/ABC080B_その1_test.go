package ABC_problemB

import (
	"testing"
)

// [ABC080B - Harshad Number](https://atcoder.jp/contests/abc080/tasks/abc080_b)
func AnswerABC080Bその1(N int) string {
	fx := 0
	tmp := N
	for tmp > 0 {
		fx += tmp % 10
		tmp /= 10
	}

	if N%fx == 0 {
		return "Yes"
	} else {
		return "No"
	}
}

func TestAnswerABC080Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want string
	}{
		{"入力例1", 12, "Yes"},
		{"入力例2", 57, "No"},
		{"入力例3", 148, "No"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC080Bその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
