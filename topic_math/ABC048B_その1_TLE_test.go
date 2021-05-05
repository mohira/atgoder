package p08_ABC085C

import (
	"testing"
)

// [ABC048B - Between a and b ... ](https://atcoder.jp/contests/abc048/tasks/abc048_b)
func AnswerABC048Bその1(a, b, x int) int {
	// 区間[a,b]で、xで割り切れる最小の数qを探す
	// ここで全探索をしているので、O(N) で 1<=b<=1e+18 なので TLEになる
	var q = -99999
	for i := a; i <= b; i++ {
		if i%x == 0 {
			q = i
			break
		}
	}

	if q < 0 {
		// q が見つからないのであれば、xで割り切れる数は1つもない
		return 0
	} else {
		// q から b までに存在する数字の個数N(qは含まない)
		N := b - q

		// x個ごとに1つの倍数が見つかるので、N個の連続する整数のなかには、 Nとxの商だけ xの倍数がある
		return 1 + N/x
	}
}

func TestAnswerABC048Bその1(t *testing.T) {
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
		//{"bとcが巨大な数の場合", 1, 1e+18, 1e+9, 1000000001},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC048Bその1(tt.a, tt.b, tt.x)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
