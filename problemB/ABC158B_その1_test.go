package problemB

import (
	"atgoder/lib"
	"testing"
)

// [ABC158B - Count Balls](https://atcoder.jp/contests/abc158/tasks/abc158_b)
func AnswerABC158Bその1(N, A, B int) int {
	var ans int
	q, r := N/(A+B), N%(A+B)

	ans += q * A
	ans += lib.Min(r, A)

	return ans
}

func TestAnswerABC158Bその1(t *testing.T) {
	tests := []struct {
		name    string
		N, A, B int
		want    int
	}{
		{"入力例1", 8, 3, 4, 4},
		{"入力例2", 8, 0, 4, 0},
		{"入力例3", 6, 2, 4, 2},

		{"入力例", 7, 3, 4, 3},
		{"入力例", 6, 3, 4, 3},
		{"入力例", 5, 3, 4, 3},
		{"入力例", 4, 3, 4, 3},
		{"入力例", 3, 3, 4, 3},
		{"入力例", 2, 3, 4, 2},
		{"入力例", 1, 3, 4, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC158Bその1(tt.N, tt.A, tt.B)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
