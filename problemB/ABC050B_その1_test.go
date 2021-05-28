package problemB

import (
	"fmt"
	"strings"
	"testing"
)

// [ABC050B - Contest with Drinks Easy](https://atcoder.jp/contests/abc050/tasks/abc050_b)
func AnswerABC050Bその1(N int, T []int, M int, PX [][]int) string {
	var out strings.Builder

	total := sumInts(T)

	for _, px := range PX {
		p, x := px[0], px[1]
		p--
		tmp := total - T[p] + x

		out.WriteString(fmt.Sprintf("%d\n", tmp))
	}

	return out.String()
}

func TestAnswerABC050Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		T    []int
		M    int
		PX   [][]int
		want string
	}{
		{
			"入力例1",
			3,
			[]int{2, 1, 4},
			4,
			[][]int{
				{1, 1},
				{2, 3},
			},
			"6\n9\n",
		},
		{
			"入力例2",
			5,
			[]int{7, 2, 3, 8, 5},
			3,
			[][]int{
				{4, 2},
				{1, 7},
				{4, 13},
			},
			"19\n25\n30\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC050Bその1(tt.N, tt.T, tt.M, tt.PX)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
