package ABC_problemC

import (
	"math"
	"testing"
)

// [ABC198C - Compass Walking](https://atcoder.jp/contests/abc198/tasks/abc198_c)
func AnswerABC198Cその1(R, X, Y int) int {
	// 原点から(X,Y)までの距離を計算して
	// 最短で距離を詰めて、最終で帳尻を合わせる
	// 除算と剰余算を使いたいのでintにしたいが、そのままintにするとずれるのでCeilを挟む
	distance := int(math.Ceil(math.Sqrt(float64(X*X + Y*Y))))

	if distance < R {
		// 必要距離がR未満の場合は必ず2手でいける(と思う)
		return 2
	}

	q := distance / R
	r := distance % R

	if r == 0 {
		return q
	} else {
		return q + 1
	}
}

func TestAnswerABC198Cその1(t *testing.T) {
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
			got := AnswerABC198Cその1(tt.R, tt.X, tt.Y)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
