package topic_dp

import (
	"testing"
)

// [EDPC G - Longest Path](https://atcoder.jp/contests/dp/tasks/dp_g)
func AnswerEDPC問題Gその1(N, M int, XY [][]int) int {
	edge := make([][]int, N+1) // 1-index
	for i := 0; i < M; i++ {
		from, to := XY[i][0], XY[i][1]
		edge[from] = append(edge[from], to)
	}

	var ans int

	memo := make([]int, N+1)
	done := make([]bool, N+1)

	for v := 1; v <= N; v++ {
		ans = Max(ans, dpE(v, edge, memo, done))
	}

	return ans
}

func dpE(v int, edge [][]int, memo []int, done []bool) int {
	ans := 0

	if done[v] {
		return memo[v]
	}

	for _, to := range edge[v] {
		ans = Max(ans, dpE(to, edge, memo, done)+1)
	}

	memo[v] = ans
	done[v] = true

	return ans
}

func TestAnswerEDPC問題Gその1(t *testing.T) {
	tests := []struct {
		name string
		N, M int
		XY   [][]int
		want int
	}{
		{
			"入力例1",
			4, 5,
			[][]int{
				{1, 2},
				{1, 3},
				{3, 2},
				{2, 4},
				{3, 4},
			},
			3,
		},
		{
			"入力例2",
			6, 3,
			[][]int{
				{2, 3},
				{4, 5},
				{5, 6},
			},
			2,
		},
		{
			"入力例3",
			5, 8,
			[][]int{
				{5, 3},
				{2, 3},
				{2, 4},
				{5, 2},
				{5, 1},
				{1, 4},
				{4, 3},
				{1, 3},
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnswerEDPC問題Gその1(tt.N, tt.M, tt.XY); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
