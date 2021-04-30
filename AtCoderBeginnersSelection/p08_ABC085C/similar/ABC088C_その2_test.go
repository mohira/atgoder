package p08_ABC088C

import (
	"testing"
)

// [ABC088C - Otoshidama](https://atcoder.jp/contests/abc088/tasks/arc086_a)
func AnswerABC088Cその2(C [3][3]int) string {
	// 解説: https://youtu.be/tXsdcYBhmtI?t=652
	var a1 = 0

	// a1が決まれば、b[i]も決まる
	b1 := C[0][0] - a1
	b2 := C[0][1] - a1
	b3 := C[0][2] - a1

	// b[1]が決まれば、a[i]も決まる
	a2 := C[1][0] - b1
	a3 := C[2][0] - b1

	// 条件を満たすかをチェックする
	cond1 := (C[0][0] == a1+b1) && (C[0][1] == a1+b2) && (C[0][2] == a1+b3)
	cond2 := (C[1][0] == a2+b1) && (C[1][1] == a2+b2) && (C[1][2] == a2+b3)
	cond3 := (C[2][0] == a3+b1) && (C[2][1] == a3+b2) && (C[2][2] == a3+b3)

	if cond1 && cond2 && cond3 {
		return "Yes"
	}
	return "No"
}

func TestAnswerABC088Cその2(t *testing.T) {
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
			got := AnswerABC088Cその2(tt.C)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
