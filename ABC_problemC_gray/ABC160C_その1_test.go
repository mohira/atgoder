package ABC_problemC_gray

import (
	"testing"
)

// [ABC159C - Tax Increase](https://atcoder.jp/contests/abc159/tasks/abc159_c)
func AnswerABC159Cその1(A, B int) int {
	// 条件に当てはまるまで全探索
	// 雑な探索の最大は A=100 のとき、100=1250*0.08
	for x := 1; x <= 1250; x++ {
		// 正確な整数切り捨てをする
		a := x * 8 / 100
		b := x * 10 / 100

		if a == A && b == B {
			return x
		}
	}
	return -1
}

func TestAnswerABC159Cその1(t *testing.T) {
	tests := []struct {
		name string
		A, B int
		want int
	}{
		{"入力例1", 2, 2, 25},
		{"入力例2", 8, 10, 100},
		{"入力例3", 19, 99, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC159Cその1(tt.A, tt.B)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
