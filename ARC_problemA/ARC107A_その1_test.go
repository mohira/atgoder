package ARC_problemA

import (
	"testing"
)

// [ARC107A - Simple Math](https://atcoder.jp/contests/arc107/tasks/arc107_a)
func AnswerARC107Aその1(A, B, C int) int {
	const X = 998244353

	sumA := (A + 1) * A / 2
	sumB := (B + 1) * B / 2
	sumC := (C + 1) * C / 2

	var ans = 1

	ans *= sumA % X
	ans *= sumB % X
	ans %= X
	ans *= sumC % X
	ans %= X

	return ans
}

func TestAnswerARC107Aその1(t *testing.T) {
	tests := []struct {
		name    string
		A, B, C int
		want    int
	}{
		{"入力例1", 1, 2, 3, 18},
		{"入力例2", 1000000000, 987654321, 123456789, 951633476},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerARC107Aその1(tt.A, tt.B, tt.C)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
