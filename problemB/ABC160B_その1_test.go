package problemB

import (
	"testing"
)

// [ABC160B - Golden Coins](https://atcoder.jp/contests/abc160/tasks/abc160_b)
func AnswerABC160Bその1(X int) int {
	// 500yen1枚のほうが5yen100枚より嬉しい
	num500yen := X / 500
	X %= 500

	num5yen := X / 5

	return 1000*num500yen + 5*num5yen
}

func TestAnswerABC160Bその1(t *testing.T) {
	tests := []struct {
		name string
		X    int
		want int
	}{
		{"入力例1", 1024, 2020},
		{"入力例2", 0, 0},
		{"入力例3", 1000000000, 2000000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC160Bその1(tt.X)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
