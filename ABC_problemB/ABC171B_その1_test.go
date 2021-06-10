package ABC_problemB

import (
	"sort"
	"testing"
)

// [ABC171B - Mix Juice](https://atcoder.jp/contests/abc171/tasks/abc171_b)
func AnswerABC171Bその1(N, K int, P []int) int {
	sort.Ints(P)

	var ans int
	for _, p := range P[:K] {
		ans += p
	}
	return ans
}

func TestAnswerABC171Bその1(t *testing.T) {
	tests := []struct {
		name string
		N, K int
		P    []int
		want int
	}{
		{"入力例1",
			5, 3,
			[]int{50, 100, 80, 120, 80},
			210,
		},
		{"入力例2",
			1, 1,
			[]int{1000},
			1000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC171Bその1(tt.N, tt.K, tt.P)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
