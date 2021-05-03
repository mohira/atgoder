package p10_ABC086C

import (
	"testing"
)

// [ABC093C - Same Integers](https://atcoder.jp/contests/abc093/tasks/arc094_a)
func AnswerABC093Cその1(A, B, C int) int {
	// 偶 奇 最大値の偶奇 最終的な値(Mとする)
	// 3  0  偶 => max(A,B,C)
	// 3  0  奇 => max(A,B,C)
	// 0  3  偶 => max(A,B,C)
	// 0  3  奇 => max(A,B,C)
	// 2  1  偶 => max(A,B,C) + 1
	// 2  1  奇 => max(A,B,C)
	// 1  2  偶 => max(A,B,C)
	// 1  2  奇 => max(A,B,C) + 1
	oddNumCount := A%2 + B%2 + C%2
	evenNumCount := (A+1)%2 + (B+1)%2 + (C+1)%2

	M := findMax(A, B, C)

	// M++ となる条件は、 `M%2 != (A+B+C)%2 != 0` で言い換えられる
	if evenNumCount == 2 && oddNumCount == 1 && M%2 == 0 {
		M++
	}

	if evenNumCount == 1 && oddNumCount == 2 && M%2 == 1 {
		M++
	}

	// 操作回数は最終的な値と現在の値とのズレ を 2で割ったもの
	totalDiff := (M - A) + (M - B) + (M - C) // `3M - sum` よりもイメージに近いので、この式を採用
	return totalDiff / 2
}

func findMax(a, b, c int) int {
	var m int

	for _, i := range []int{a, b, c} {
		if m < i {
			m = i
		}
	}
	return m
}

func TestAnswerABC093Cその1(t *testing.T) {
	tests := []struct {
		name string
		A    int
		B    int
		C    int
		want int
	}{
		// 偶2 奇1
		{"入力例1 偶2 奇1 max=奇数", 2, 4, 5, 2},
		{"入力例2 偶2 奇1 max=偶数", 2, 6, 3, 5},

		// 偶1 奇2
		{"入力例 偶1 奇2 max=奇数", 2, 5, 5, 3},
		{"入力例 偶1 奇2 max=偶数", 6, 5, 5, 1},

		// 偶3 奇0
		{"入力例 偶3 奇0 値が異なる", 2, 4, 6, 3},
		{"入力例 偶3 奇0 値がすべて一致", 2, 2, 2, 0},

		// 偶0 奇3
		{"入力例 偶0 奇3 値が異なる", 1, 3, 5, 3},
		{"入力例 偶0 奇3 値がすべて一致", 1, 1, 1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC093Cその1(tt.A, tt.B, tt.C)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
