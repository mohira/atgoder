package problemB

import (
	"atgoder/lib"
	"math"
	"testing"
)

// [ABC180B - B Various distances](https://atcoder.jp/contests/abc180/tasks/abc180_b)
func AnswerABC180Bその1(N int, X []int) (int, float64, int) {
	return manhattanDistance(X), euclideanDistance(X), chebyshevDistance(X)
}

func chebyshevDistance(nums []int) int {
	var distance int
	for _, num := range nums {
		distance = lib.Max(distance, lib.AbsInt(num))
	}

	return distance
}

func euclideanDistance(nums []int) float64 {
	var squaredSum int

	for _, num := range nums {
		squaredSum += num * num
	}

	return math.Sqrt(float64(squaredSum))
}

func manhattanDistance(nums []int) int {
	var distance int

	for _, num := range nums {
		distance += lib.AbsInt(num)
	}

	return distance
}

func TestAnswerABC180Bその1(t *testing.T) {
	tests := []struct {
		name  string
		N     int
		X     []int
		want1 int
		want2 float64
		want3 int
	}{
		{"入力例1",
			2,
			[]int{2, -1},
			3,
			2.236067977499790,
			2,
		},
		{"入力例2",
			10,
			[]int{3, -1, -4, 1, -5, 9, 2, -6, 5, -3},
			39,
			14.387494569938159,
			9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2, got3 := AnswerABC180Bその1(tt.N, tt.X)

			if got1 != tt.want1 {
				t.Errorf("got %v want %v", got1, tt.want1)
			}
			if math.Abs(got2-tt.want2) > 1e-9 {
				t.Errorf("got %v want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("got %v want %v", got3, tt.want3)
			}
		})
	}
}
