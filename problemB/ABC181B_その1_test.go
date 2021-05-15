package problemB

import (
	"testing"
)

// [ABC181B - Trapezoid Sum](https://atcoder.jp/contests/abc181/tasks/abc181_b)
func AnswerABC181Bその1(N int, AB [][]int) int {
	var sum int

	for i := 0; i < N; i++ {
		a, b := AB[i][0], AB[i][1]

		sum += sumIntRange(a, b)
	}

	return sum
}

func TestAnswerABC181Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		AB   [][]int
		want int
	}{
		{"入力例1",
			2,
			[][]int{
				{1, 3},
				{3, 5},
			},
			18},
		{"入力例2",
			3,
			[][]int{
				{11, 13},
				{17, 47},
				{359, 44683},
			},
			998244353,
		},
		{"入力例3",
			1,
			[][]int{
				{1, 1000000},
			},
			500000500000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC181Bその1(tt.N, tt.AB)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
