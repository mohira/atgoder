package similar

import (
	"math"
	"testing"
)

// [ABC102B - Maximum Difference](https://atcoder.jp/contests/abc102/tasks/abc102_b)
func AnswerABC102Bその1(N int, A []int) int {
	// 愚直な二重ループ
	var maxDiff int

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			diff := int(math.Abs(float64(A[i] - A[j])))

			if maxDiff < diff {
				maxDiff = diff
			}
		}
	}

	return maxDiff
}

func TestAnswerABC102Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want int
	}{
		{"入力例1", 4, []int{1, 4, 6, 3}, 5},
		{"入力例2", 2, []int{1000000000, 1}, 999999999},
		{"入力例3", 5, []int{1, 1, 1, 1, 1}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC102Bその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}

		})

	}
}
