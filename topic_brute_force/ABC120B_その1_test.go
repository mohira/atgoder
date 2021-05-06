package topic_brute_force

import (
	"testing"
)

// [ABC120B - 105](https://atcoder.jp/contests/abc120/tasks/abc120_b)
func AnswerABC120Bその1(A, B, K int) int {
	var candidates = make([]int, 0)

	for i := 1; i <= 100; i++ {
		if A%i == 0 && B%i == 0 {
			candidates = append(candidates, i)
		}
	}

	return candidates[len(candidates)-K]

	// ソートした方がindex間違えにくそう
	// sort.Sort(sort.Reverse(sort.IntSlice(candidates)))
	// return candidates[K-1]
}

func TestAnswerABC120Bその1(t *testing.T) {
	tests := []struct {
		name string
		A    int
		B    int
		K    int
		want int
	}{
		{"入力例1", 8, 12, 2, 2},
		{"入力例2", 100, 50, 4, 5},
		{"入力例3", 1, 1, 1, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC120Bその1(tt.A, tt.B, tt.K)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
