package topic_counter

import (
	"testing"
)

func AnswerABC202C(N int, A, B, C []int) int {
	return -1
}

func TestABC202C(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		B    []int
		C    []int
		want int
	}{
		{
			"入力例1",
			3,
			[]int{1, 2, 2},
			[]int{3, 1, 2},
			[]int{2, 3, 2},
			4,
		},
		{
			"入力例2",
			4,
			[]int{1, 1, 1, 1},
			[]int{1, 1, 1, 1},
			[]int{1, 2, 3, 4},
			16,
		},
		{
			"入力例3",
			3,
			[]int{2, 3, 3},
			[]int{1, 3, 3},
			[]int{1, 1, 1},
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC202C(tt.N, tt.A, tt.B, tt.C)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}

}
