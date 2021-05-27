package problemB

import (
	"testing"
)

// [ABC092B - Chocolate](https://atcoder.jp/contests/abc092/tasks/abc092_b)
func AnswerABC092Bその1(N, D, X int, A []int) int {
	count := 0
	for _, a := range A {
		distributedDay := 1

		// (D-1)/ a + 1 でも計算できる
		for distributedDay <= D {
			distributedDay += a
			count++
		}
	}

	return count + X
}

func TestAnswerABC092Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		D, X int
		A    []int
		want int
	}{
		{
			"入力例1",
			3, 7, 1,
			[]int{2, 5, 10},
			8,
		},
		{
			"入力例2",
			2, 8, 20,
			[]int{1, 10},
			29,
		},
		{
			"入力例3",
			5, 30, 44,
			[]int{26, 18, 81, 18, 6},
			56,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC092Bその1(tt.N, tt.D, tt.X, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
