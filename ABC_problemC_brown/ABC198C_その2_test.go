package ABC_problemC_brown

import (
	"math"
	"testing"
)

// [ABC198C - Compass Walking](https://atcoder.jp/contests/abc198/tasks/abc198_c)
func AnswerABC198Cその2(R, X, Y int) int {
	distance := math.Sqrt(float64(X*X + Y*Y))

	switch {
	case int(distance) == R:
		// ぴったりなら一撃
		return 1
	case distance < float64(R):
		// R未満なら2手でいける
		return 2
	default:
		// 直線で向かってラスト2手で帳尻合わせ
		// 剰余算使わなくてもCeilを使えばfloatのまま処理できる
		return int(math.Ceil(distance / float64(R)))
	}
}

func TestAnswerABC198Cその2(t *testing.T) {
	tests := []struct {
		name    string
		R, X, Y int
		want    int
	}{
		{"入力例1", 5, 15, 0, 3},
		{"入力例2", 5, 11, 0, 3},
		{"入力例3", 3, 4, 4, 2},
		{"", 95044, 28016, 39332, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC198Cその2(tt.R, tt.X, tt.Y)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
