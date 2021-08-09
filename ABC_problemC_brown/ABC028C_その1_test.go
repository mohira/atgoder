package ABC_problemC_brown

import (
	"sort"
	"testing"
)

// [ABC028C - 数を3つ選ぶマン](https://atcoder.jp/contests/abc028/tasks/abc028_c)
func AnswerABC028Cその1(A, B, C, D, E int) int {
	x := []int{A, B, C, D, E}

	nums := make([]int, 0, 10)

	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			for k := j + 1; k < 5; k++ {
				nums = append(nums, x[i]+x[j]+x[k])
			}
		}
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	return nums[7]
}

func TestAnswerABC028Cその1(t *testing.T) {
	tests := []struct {
		name          string
		A, B, C, D, E int
		want          int
	}{
		{"入力例1", 1, 2, 3, 4, 5, 10},
		{"入力例2", 1, 2, 3, 5, 8, 14},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC028Cその1(tt.A, tt.B, tt.C, tt.D, tt.E)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
