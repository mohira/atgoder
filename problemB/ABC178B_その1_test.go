package problemB

import (
	"testing"
)

// [ABC178B - Product Max](https://atcoder.jp/contests/abc178/tasks/abc178_b)
func AnswerABC178Bその1(a, b, c, d int) int {
	// 1行でも書ける
	// return max(max(a*c, a*d), max(b*c, b*d))

	xs := []int{a, b}
	ys := []int{c, d}

	var productMax int = -1e+18

	for _, x := range xs {
		for _, y := range ys {
			product := x * y

			productMax = max(productMax, product)
		}
	}

	return productMax
}

func TestAnswerABC178Bその1(t *testing.T) {
	tests := []struct {
		name       string
		a, b, c, d int
		want       int
	}{
		{"入力例1", 1, 2, 1, 1, 2},
		{"入力例2", 3, 5, -4, -2, -6},
		{"入力例3", -1000000000, 0, -1000000000, 0, 1000000000000000000},
		{"でかいケース", -910805388, -6344404, 430910689, 796838524, -2733871498934356},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC178Bその1(tt.a, tt.b, tt.c, tt.d)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
