package ABC_problemA

import (
	"testing"
)

// [ABC022A - Best Body](https://atcoder.jp/contests/abc022/tasks/abc022_a)
func AnswerABC022Aその1(N, S, T, W int, A []int) int {
	count := 0

	if S <= W && W <= T {
		count++
	}

	for _, a := range A {
		W += a
		if S <= W && W <= T {
			count++
		}
	}

	return count
}

func TestAnswerABC022Aその1(t *testing.T) {
	tests := []struct {
		name    string
		N, S, T int
		W       int
		A       []int
		want    int
	}{
		{
			"入力例1",
			5, 60, 70,
			50,
			[]int{10, 10, 10, 10},
			2,
		},
		{
			"入力例2",
			5, 50, 100,
			120,
			[]int{-10, -20, -30, 70},
			2,
		},
		{
			"初日からベストボディ",
			1, 50, 100,
			50,
			[]int{1000},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC022Aその1(tt.N, tt.S, tt.T, tt.W, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
