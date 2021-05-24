package problemB

import (
	"testing"
)

// [ABC121B - Can you solve this?](https://atcoder.jp/contests/abc121/tasks/abc121_b)
func AnswerABC121Bその1(N, M, C int, B []int, A [][]int) int {
	var count int
	for i := 0; i < N; i++ {
		feature := 0

		for j := 0; j < M; j++ {
			feature += A[i][j] * B[j]
		}

		feature += C
		if feature > 0 {
			count++
		}
	}
	return count
}

func TestAnswerABC121Bその1(t *testing.T) {
	tests := []struct {
		name    string
		N, M, C int
		B       []int
		A       [][]int
		want    int
	}{
		{
			"入力例1",
			2, 3, -10,
			[]int{1, 2, 3},
			[][]int{
				{3, 2, 1},
				{1, 2, 2},
			},
			1,
		},
		{
			"入力例2",
			5, 2, -4,
			[]int{-2, 5},
			[][]int{
				{100, 41},
				{100, 40},
				{-3, 0},
				{-6, -2},
				{18, -13},
			},
			2,
		},
		{
			"入力例3",
			3, 3, 0,
			[]int{100, -100, 0},
			[][]int{
				{0, 100, 100},
				{100, 100, 100},
				{-100, 100, 100},
			},
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC121Bその1(tt.N, tt.M, tt.C, tt.B, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
