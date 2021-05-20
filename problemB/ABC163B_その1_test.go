package problemB

import (
	"testing"
)

// [ABC163B - Homework](https://atcoder.jp/contests/abc163/tasks/abc163_b)
func AnswerABC163Bその1(N, M int, A []int) int {
	var requireDays int
	for _, a := range A {
		requireDays += a
	}

	if requireDays <= N {
		return N - requireDays
	} else {
		return -1
	}
}

func TestAnswerABC163Bその1(t *testing.T) {
	tests := []struct {
		name string
		N, M int
		A    []int
		want int
	}{
		{"入力例1",
			41, 2,
			[]int{5, 6},
			30,
		},
		{"入力例2",
			10, 2,
			[]int{5, 6},
			-1,
		},
		{"入力例3",
			11, 2,
			[]int{5, 6},
			0,
		},
		{"入力例4",
			314, 5,
			[]int{9, 26, 5, 35, 8, 9, 79, 3, 23, 8, 46, 2, 6, 43, 3},
			9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC163Bその1(tt.N, tt.M, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
