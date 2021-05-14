package problemB

import (
	"testing"
)

// [ABC186B - Gentle Pairs](https://atcoder.jp/contests/abc186/tasks/abc186_b)
func AnswerABC186Bその1(H, W int, A [][]int) int {
	var minBlock = 100
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			minBlock = min(minBlock, A[i][j])
		}
	}

	var ans int
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			ans += A[i][j] - minBlock
		}
	}

	return ans
}

func TestAnswerABC186Bその1(t *testing.T) {
	tests := []struct {
		name string
		H, W int
		A    [][]int
		want int
	}{
		{"入力例1",
			2, 3,
			[][]int{
				{2, 2, 3},
				{3, 2, 2},
			},
			2,
		},
		{"入力例2",
			3, 3,
			[][]int{
				{99, 99, 99},
				{99, 0, 99},
				{99, 99, 99},
			},
			792,
		},
		{"入力例3",
			3, 2,
			[][]int{
				{4, 4},
				{4, 4},
				{4, 4},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC186Bその1(tt.H, tt.W, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
