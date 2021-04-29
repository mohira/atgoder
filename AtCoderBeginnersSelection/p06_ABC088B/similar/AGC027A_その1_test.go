package similar

import (
	"sort"
	"testing"
)

// [A - Candy Distribution Again](https://atcoder.jp/contests/agc027/tasks/agc027_a)
func AnswerAGC027Aその1(N, x int, A []int) int {
	var ans int

	sort.Ints(A)

	for i := 0; i < N; i++ {
		a := A[i]
		if i == N-1 {
			// 最後の子供のチェックはぴったり判定が必要
			if x == a {
				ans++
			}
		} else if x-a >= 0 {
			ans++
			x -= a
		}
	}
	return ans
}

func TestAnswerAGC027Aその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		x    int
		A    []int
		want int
	}{
		{"入力例1", 3, 70, []int{20, 30, 10}, 2},
		{"入力例2", 3, 10, []int{20, 30, 10}, 1},
		{"入力例3", 4, 1111, []int{1, 10, 100, 1000}, 4},
		{"入力例4", 2, 10, []int{20, 20}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerAGC027Aその1(tt.N, tt.x, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
