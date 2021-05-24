package problemB

import (
	"testing"
)

// [ABC139B - Power Socket](https://atcoder.jp/contests/abc139/tasks/abc139_b)
func AnswerABC139Bその1(A, B int) int {
	var taps = 1
	var ans int

	for {
		if B <= taps {
			return ans
		}

		taps += A - 1
		ans++
	}
}

func TestAnswerABC139Bその1(t *testing.T) {
	tests := []struct {
		name string
		A, B int
		want int
	}{
		{"入力例1", 4, 10, 3},
		{"入力例2", 8, 9, 2},
		{"入力例3", 8, 8, 1},
		{"追加のタップが不要なケース", 1, 1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC139Bその1(tt.A, tt.B)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
