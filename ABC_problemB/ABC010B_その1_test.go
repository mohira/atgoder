package ABC_problemB

import (
	"atgoder/lib"
	"testing"
)

// [ABC010B - 花占い](https://atcoder.jp/contests/abc010/tasks/abc010_2)
func AnswerABC010Bその1(N int, A []int) int {
	var ans int

	oks := []int{1, 3, 7, 9}

	const INF = 10

	for _, a := range A {
		min := INF

		for _, ok := range oks {
			if ok <= a {
				min = lib.Min(min, a-ok)
			}
		}

		ans += min
	}

	return ans
}

func TestAnswerABC010Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want int
	}{
		{"入力例1", 3, []int{5, 8, 2}, 4},
		{"入力例2", 9, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 8},
		{"入力例3", 3, []int{3, 1, 4}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC010Bその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
