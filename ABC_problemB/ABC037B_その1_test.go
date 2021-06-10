package ABC_problemB

import (
	"bytes"
	"strconv"
	"testing"
)

// [ABC037B - 編集](https://atcoder.jp/contests/abc037/tasks/abc037_b)
func AnswerABC037Bその1(N, Q int, LRT [][]int) string {
	nums := make([]int, N)

	for _, lrt := range LRT {
		l, r, t := lrt[0], lrt[1], lrt[2]
		l--
		r--
		for i := l; i <= r; i++ {
			nums[i] = t
		}
	}

	var out bytes.Buffer
	for _, num := range nums {
		out.WriteString(strconv.Itoa(num) + "\n")
	}
	return out.String()
}

func TestAnswerABC037Bその1(t *testing.T) {
	tests := []struct {
		name string
		N, Q int
		LRT  [][]int
		want string
	}{
		{
			"入力例1",
			5, 2,
			[][]int{
				{1, 3, 10},
				{2, 4, 20},
			},
			"10\n20\n20\n20\n0\n",
		},
		{
			"入力例2",
			10, 4,
			[][]int{
				{2, 7, 22},
				{3, 5, 4},
				{6, 10, 1},
				{4, 4, 12},
			},
			"0\n22\n4\n12\n4\n1\n1\n1\n1\n1\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC037Bその1(tt.N, tt.Q, tt.LRT)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
