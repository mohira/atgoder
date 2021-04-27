package similar

import (
	"testing"
)

// [ABC102B - Maximum Difference](https://atcoder.jp/contests/abc102/tasks/abc102_b)
func AnswerABC102Bその2(N int, A []int) int {
	return MaxInt(A) - MinInt(A)
}

func MaxInt(nums []int) int {
	var max = 1

	for _, n := range nums {
		if max <= n {
			max = n
		}
	}

	return max
}

func MinInt(nums []int) int {
	var min = 1_000_000_000

	for _, n := range nums {
		if n < min {
			min = n
		}
	}

	return min
}

func TestAnswerABC102Bその2(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want int
	}{
		{"入力例1", 4, []int{1, 4, 6, 3}, 5},
		{"入力例2", 2, []int{1000000000, 1}, 999999999},
		{"入力例3", 5, []int{1, 1, 1, 1, 1}, 0},
		{"入力例4", 2, []int{877914575, 602436426}, 275478149},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC102Bその2(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}

		})

	}
}
