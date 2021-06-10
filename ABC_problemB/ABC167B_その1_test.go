package ABC_problemB

import (
	"testing"
)

// [ABC167B - Easy Linear Programming](https://atcoder.jp/contests/abc167/tasks/abc167_b)
func AnswerABC167Bその1(A, B, C, K int) int {
	var point int

	// Aだけで満たせるならそれでおしまい
	if K <= A {
		return +1 * K
	}

	// Aをとれるだけとる
	point += +1 * A
	K -= A

	if K <= 0 {
		return point
	}

	// Bをとれるだけとってみる
	K -= B
	if K <= 0 {
		return point
	}

	// 残り枚数分だけ引く
	point += -1 * K

	return point
}

func TestAnswerABC167Bその1(t *testing.T) {
	tests := []struct {
		name       string
		A, B, C, K int
		want       int
	}{
		{"入力例1", 2, 1, 1, 3, 2},
		{"入力例2", 1, 2, 3, 4, 0},
		{"入力例3", 2000000000, 0, 0, 2000000000, 2000000000},
		{"入力例4", 1000000000, 500000000, 500000000, 1, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC167Bその1(tt.A, tt.B, tt.C, tt.K)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
