package topic_grid

import (
	"testing"
)

// [ABC070B - Two Switches](https://atcoder.jp/contests/abc070/tasks/abc070_b)
func AnswerABC070Bその1(A, B, C, D int) int {
	// 重ならないのはこの2パターン
	// A    B   C     D
	// l----u
	//          l-----u
	//
	// C    D   A     B
	// l----u
	//          l-----u
	maxLower := max(A, C)
	minUpper := min(B, D)

	notIntersection := minUpper < maxLower
	if notIntersection {
		return 0
	} else {
		return minUpper - maxLower
	}

	// maxを使って書いてもいいと思う
	// return max(minUpper-maxLower, 0)
}

func max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func TestAnswerABC070Bその1(t *testing.T) {
	tests := []struct {
		name string
		A    int
		B    int
		C    int
		D    int
		want int
	}{
		{"入力例1: A<C<B<D 部分重なり", 0, 75, 25, 100, 50},
		{"入力例2: A<B<C<D 重なりナシ", 0, 33, 66, 99, 0},
		{"入力例3: A<C<D<B 包含", 10, 90, 20, 80, 60},

		//
		{"入力例4: C<A<D<B 部分重なり", 25, 100, 0, 75, 50},
		{"入力例5: C<D<A<B 重なりナシ", 66, 99, 0, 33, 0},
		{"入力例6: C<A<B<D 包含", 20, 80, 10, 90, 60},

		// 境界値
		{"入力例7", 0, 75, 0, 75, 75},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC070Bその1(tt.A, tt.B, tt.C, tt.D)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
