package others

import (
	"strings"
	"testing"
)

func MSolutions2020AnswerC(N, K int, A []int) string {
	var out = make([]string, N-K)

	for i := K; i < N; i++ {
		left, right := A[i-K], A[i]

		if left < right {
			out[i-K] = "Yes"
		} else {
			out[i-K] = "No"
		}
	}

	return strings.Join(out, "\n")
}

func TestMSolutions2020AnswerC(t *testing.T) {
	tests := []struct {
		name string
		N, K int
		A    []int
		want string
	}{
		{
			"入力例1",
			5, 3,
			[]int{96, 98, 95, 100, 20},
			"Yes\nNo",
		},
		{
			"入力例2",
			3, 2,
			[]int{1001, 869120, 1001},
			"No",
		},
		{
			"入力例3",
			15, 7,
			[]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9},
			"Yes\nYes\nNo\nYes\nYes\nNo\nYes\nYes",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MSolutions2020AnswerC(tt.N, tt.K, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
