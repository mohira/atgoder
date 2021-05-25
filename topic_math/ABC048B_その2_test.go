package topic_math

import (
	"testing"
)

// [ABC048B - Between a and b ... ](https://atcoder.jp/contests/abc048/tasks/abc048_b)
func AnswerABC048Bその2(a, b, x int) int {
	// 0 x x x a o o o o o o b
	// |---------------------|
	// |-----| |<----------->|
	//            求めたい区間
	Nb := f(b, x)
	Na := f(a-1, x)

	return Nb - Na
}

func f(n, x int) int {
	// 0以上n以下の整数のうち、xで割り切れる数字の個数
	// 0は確定で割れるので +1
	// 1以上では、「x個 ごとに xの倍数が出現する」
	if n >= 0 {
		return 1 + (n / x)
	} else {
		return 0
	}
}

func TestAnswerABC048Bその2(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		x    int
		want int
	}{
		{"入力例1", 4, 8, 2, 3},
		{"入力例2", 0, 5, 1, 6},
		{"入力例3", 9, 9, 2, 0},
		{"入力例4", 1, 1e+18, 3, 333333333333333333},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC048Bその2(tt.a, tt.b, tt.x)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
