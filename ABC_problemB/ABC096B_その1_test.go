package ABC_problemB

import (
	"atgoder/lib"
	"testing"
)

// [ABC096B - Maximum Sum](https://atcoder.jp/contests/abc096/tasks/abc096_b)
func AnswerABC096Bその1(A, B, C, K int) int {
	nums := []int{A, B, C}
	maxV := lib.FindMax(nums)
	total := lib.SumInts(nums)

	double := maxV
	for i := 0; i < K; i++ {
		double *= 2
	}
	return total - maxV + double
}

func TestAnswerABC096Bその1(t *testing.T) {
	tests := []struct {
		name    string
		A, B, C int
		K       int
		want    int
	}{
		{"入力例1", 5, 3, 11, 1, 30},
		{"入力例2", 3, 3, 4, 2, 22},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC096Bその1(tt.A, tt.B, tt.C, tt.K)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
