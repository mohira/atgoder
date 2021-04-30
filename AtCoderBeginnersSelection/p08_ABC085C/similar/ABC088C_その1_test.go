package p08_ABC085C

import (
	"testing"
)

// [ABC088C - Takahashi's Information](https://atcoder.jp/contests/abc088/tasks/arc086_a)
func AnswerABC088Cその1(C [3][3]int) string {
	// 探索範囲を絞って全探索
	c11, c22, c33 := C[0][0], C[1][1], C[2][2]

	for a1 := 0; a1 <= c11; a1++ {
		for a2 := 0; a2 <= c22; a2++ {
			for a3 := 0; a3 <= c33; a3++ {
				b1 := c11 - a1 // ⇔ a1 + b1 = c11
				b2 := c22 - a2 // ⇔ a2 + b2 = c22
				b3 := c33 - a3 // ⇔ a3 + b3 = c33

				if correctPrediction(C, a1, a2, a3, b1, b2, b3) {
					return "Yes"
				}
			}
		}
	}

	return "No"
}

func correctPrediction(C [3][3]int, a1, a2, a3, b1, b2, b3 int) bool {
	c11, c12, c13 := C[0][0], C[0][1], C[0][2]
	c21, c22, c23 := C[1][0], C[1][1], C[1][2]
	c31, c32, c33 := C[2][0], C[2][1], C[2][2]

	matchLine1 := (c11 == a1+b1) && (c12 == a1+b2) && (c13 == a1+b3)
	matchLine2 := (c21 == a2+b1) && (c22 == a2+b2) && (c23 == a2+b3)
	matchLine3 := (c31 == a3+b1) && (c32 == a3+b2) && (c33 == a3+b3)

	return matchLine1 && matchLine2 && matchLine3
}

func TestAnswerABC088Cその1(t *testing.T) {
	tests := []struct {
		name string
		C    [3][3]int
		want string
	}{
		{"入力例1", [3][3]int{
			{1, 0, 1},
			{2, 1, 2},
			{1, 0, 1},
		}, "Yes"},
		{"入力例2", [3][3]int{
			{2, 2, 2},
			{2, 1, 2},
			{2, 2, 2},
		}, "No"},
		{"入力例3", [3][3]int{
			{0, 8, 8},
			{0, 8, 8},
			{0, 8, 8},
		}, "Yes"},

		{"入力例4", [3][3]int{
			{1, 8, 6},
			{2, 9, 7},
			{0, 7, 7},
		}, "No"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC088Cその1(tt.C)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
