package topic_cumlative_sum

import (
	"testing"
)

// [ABC127C - Prison](https://atcoder.jp/contests/abc127/tasks/abc127_c)
func AnswerABC127Cその1(N, M int, LR [][]int) int {
	imos := make([]int, N+1)

	for i := 0; i < M; i++ {
		l, r := LR[i][0], LR[i][1]
		l--
		r--

		imos[l]++
		imos[r+1]--
	}

	cumsum := getCumulativeSum(imos)

	var count int
	for _, s := range cumsum[1 : len(cumsum)-1] {
		if s == M {
			count++
		}
	}
	return count
}

func TestAnswerABC127Cその1(t *testing.T) {
	tests := []struct {
		name string
		N, M int
		LR   [][]int
		want int
	}{
		{"入力例1",
			4, 2,
			[][]int{
				{1, 3},
				{2, 4},
			},
			2},
		{"入力例2",
			10, 3,
			[][]int{
				{3, 6},
				{5, 7},
				{6, 9},
			},
			1},
		{"入力例3",
			100000, 1,
			[][]int{
				{1, 100000},
			},
			100000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC127Cその1(tt.N, tt.M, tt.LR)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
