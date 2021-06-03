package problemC

import (
	"atgoder/lib"
	"testing"
)

// [ABC135C - City Savers](https://atcoder.jp/contests/abc135/tasks/abc135_c)
func AnswerABC135Cその1(N int, A, B []int) int {
	count := 0

	for i := 0; i < N; i++ {
		diff1 := lib.Min(B[i], A[i])
		B[i] -= diff1
		A[i] -= diff1
		count += diff1

		diff2 := lib.Min(B[i], A[i+1])
		B[i] -= diff2
		A[i+1] -= diff2
		count += diff2
	}

	return count
}

func TestAnswerABC135Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A, B []int
		want int
	}{
		{
			"入力例1",
			2,
			[]int{3, 5, 2},
			[]int{4, 5},
			9,
		},
		{
			"入力例2",
			3,
			[]int{5, 6, 3, 8},
			[]int{5, 100, 8},
			22,
		},
		{
			"入力例3",
			2,
			[]int{100, 1, 1},
			[]int{1, 100},
			3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC135Cその1(tt.N, tt.A, tt.B)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
