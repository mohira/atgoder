package similar

import (
	"testing"
)

// [ABC088A - Infinite Coins](https://atcoder.jp/contests/abc088/tasks/abc088_a)
func AnswerABC088A(N int, A int) string {
	// 500円を使い切る → 残りがA円以下ならぴったり支払いが可能
	if N%500 <= A {
		return "Yes"
	} else {
		return "No"
	}
}

func TestAnswerABC088A(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    int
		want string
	}{
		{"入力例1", 2018, 218, "Yes"},
		{"入力例2", 2763, 0, "No"},
		{"入力例3", 37, 514, "Yes"},

		{"入力例X", 2763, 9999999999999, "Yes"},
		{"入力例X", 500, 0, "Yes"},
		{"入力例X", 499, 0, "No"},
		{"入力例X", 499, 500, "Yes"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC088A(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
