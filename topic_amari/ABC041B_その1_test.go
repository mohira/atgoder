package topic_amari

import (
	"testing"
)

// [ABC041B - 直方体](https://atcoder.jp/contests/abc041/tasks/abc041_b)
func AnswerABC041Bその1(A, B, C int) int {
	const X int = 1e+9 + 7

	return A * B % X * C % X
}

func TestAnswerABC041Bその1(t *testing.T) {
	tests := []struct {
		name    string
		A, B, C int
		want    int
	}{
		{"入力例1", 2, 3, 4, 24},
		{"入力例2", 10000, 1000, 100, 1000000000},
		{"入力例3", 100000, 1, 100000, 999999937},
		{"入力例4", 1000000000, 1000000000, 1000000000, 999999664},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC041Bその1(tt.A, tt.B, tt.C)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
