package p10_ABC086C

import (
	"testing"
)

// [ABC093C - Same Integers](https://atcoder.jp/contests/abc093/tasks/arc094_a)
func AnswerABC093Cその2(A, B, C int) int {
	// 3つの整数の和 の偶奇 は 操作によって不変(+2ずつしか変化しないからね)
	// M = max(A, B, C) とする
	// 3M と 3つの整数の和 が一致      => A,B,Cは M   に向かえばいい {A, B, C} -> {M,   M,   M}
	// 3M と 3つの整数の和 が一致しない => A,B,Cは M+1 に向かえばいい {A, B, C} -> {M+1, M+1, M+1}
	M := findMax(A, B, C)
	sum := A + B + C

	偶奇が一致 := 3*M%2 == sum%2

	if 偶奇が一致 {
		totalDiff := (M - A) + (M - B) + (M - C)
		return totalDiff / 2
	} else {
		totalDiff := (M + 1 - A) + (M + 1 - B) + (M + 1) - C
		return totalDiff / 2
	}
}

func TestAnswerABC093Cその2(t *testing.T) {
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
			got := AnswerABC093Cその2(tt.A, tt.B, tt.C)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
