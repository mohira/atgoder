package topic_brute_force

import (
	"testing"
)

// [ABC120B - 105](https://atcoder.jp/contests/abc120/tasks/abc120_b)
func AnswerABC120Bその2(A, B, K int) int {
	// 減少させていくやり方面白い
	for i := maxInt(A, B); i >= 1; i-- {
		if A%i == 0 && B%i == 0 {
			K--

			if K == 0 {
				return i
			}
		}
	}

	return 0
}

func maxInt(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func TestAnswerABC120Bその2(t *testing.T) {
	tests := []struct {
		name string
		A    int
		B    int
		K    int
		want int
	}{
		{"入力例1", 8, 12, 2, 2},
		{"入力例2", 100, 50, 4, 5},
		{"入力例3", 1, 1, 1, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC120Bその2(tt.A, tt.B, tt.K)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
