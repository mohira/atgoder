package problemB

import (
	"sort"
	"testing"
)

// [ABC064B - Traveling AtCoDeer Problem](https://atcoder.jp/contests/abc077/tasks/abc077_b)
func AnswerABC064Bその1(N int, A []int) int {
	sort.Ints(A)

	var ans int

	for i := 0; i < N-1; i++ {
		ans += A[i+1] - A[i]
	}

	return ans
}

func TestAnswerABC064Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want int
	}{
		{"入力例1", 4, []int{2, 3, 7, 9}, 7},
		{"入力例2", 8, []int{3, 1, 4, 1, 5, 9, 2, 6}, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC064Bその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
