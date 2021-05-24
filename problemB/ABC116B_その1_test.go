package problemB

import (
	"testing"
)

// [ABC1166B - Collatz Problem](https://atcoder.jp/contests/abc1166/tasks/abc1166_b)
func AnswerABC1166Bその1(s int) int {
	bucket := make(map[int]int)
	ai := s
	bucket[ai]++

	// 条件を満たす最小のmは1000000以下となることが保証される。
	for m := 2; m <= 1000000; m++ {
		ai = f(ai)
		if _, ok := bucket[ai]; ok {
			return m
		} else {
			bucket[ai]++
		}
	}
	return 0
}

func f(n int) int {
	if n%2 == 0 {
		return n / 2
	} else {
		return 3*n + 1
	}
}

func TestAnswerABC1166Bその1(t *testing.T) {
	tests := []struct {
		name string
		s    int
		want int
	}{
		{"入力例1", 8, 5},
		{"入力例2", 7, 18},
		{"入力例3", 54, 114},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC1166Bその1(tt.s)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
