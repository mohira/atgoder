package ABC_problemB

import (
	"atgoder/lib"
	"testing"
)

// [ABC199B - Intersection](https://atcoder.jp/contests/abc199/tasks/abc199_b)
func AnswerABC199Bその1(N int, A []int, B []int) int {
	maxA := lib.FindMax(A)
	minB := lib.FindMin(B)

	if minB-maxA >= 0 {
		return minB - maxA + 1
	} else {
		return 0
	}
}

func TestAnswerABC199Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		B    []int
		want int
	}{
		{"入力例1",
			2,
			[]int{3, 2},
			[]int{7, 5},
			3,
		},
		{"入力例2",
			3,
			[]int{1, 5, 3},
			[]int{10, 7, 3},
			0,
		},
		{"入力例3",
			3,
			[]int{3, 2, 5},
			[]int{6, 9, 8},
			2,
		},
		{"境界条件",
			1,
			[]int{2},
			[]int{2},
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC199Bその1(tt.N, tt.A, tt.B)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
