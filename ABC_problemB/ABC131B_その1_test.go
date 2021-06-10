package ABC_problemB

import (
	"atgoder/lib"
	"math"
	"testing"
)

// [ABC131B - Bite Eating](https://atcoder.jp/contests/abc131/tasks/abc131_b)
func AnswerABC131Bその1(N, L int) int {
	var totalFlavor int
	for i := 1; i <= N; i++ {
		flavor := L + i - 1
		totalFlavor += flavor
	}

	var ans int
	var minDiff = math.MaxInt64
	for i := 1; i <= N; i++ {
		flavor := L + i - 1
		diff := lib.AbsInt(flavor)

		if diff < minDiff {
			ans = totalFlavor - flavor
			minDiff = diff
		}
	}

	return ans
}

func TestAnswerABC131Bその1(t *testing.T) {
	tests := []struct {
		name string
		N, L int
		want int
	}{
		{"入力例1", 5, 2, 18},
		{"入力例2", 3, -1, 0},
		{"入力例3", 30, -50, -1044},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC131Bその1(tt.N, tt.L)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
