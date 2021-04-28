package similar

import (
	"testing"
)

// [ABC095B - Bitter Alchemy](https://atcoder.jp/contests/abc095/tasks/abc095_b)
func AnswerABC095Bその1(N int, X int, M []int) int {
	available := X - mySum(M)

	return N + available/myMin(M)
}

func myMin(nums []int) int {
	var min = 100_000

	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	return min
}

func mySum(nums []int) int {
	var sum int

	for _, num := range nums {
		sum += num
	}

	return sum
}

func TestAnswerABC095Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		X    int
		M    []int
		want int
	}{
		{"入力例1", 3, 1000, []int{120, 100, 140}, 9},
		{"入力例2", 4, 360, []int{90, 90, 90, 90}, 4},
		{"入力例3", 5, 3000, []int{150, 130, 150, 130, 110}, 26},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC095Bその1(tt.N, tt.X, tt.M)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
