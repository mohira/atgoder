package topic_cumlative_sum

import (
	"testing"
)

// [ABC014C - AtColor](https://atcoder.jp/contests/abc014/tasks/abc014_3)
func AnswerABC014Cその1(N int, AB [][]int) int {
	// いもす法
	var S = make([]int, (1e+6)+2)
	for _, ab := range AB {
		a, b := ab[0], ab[1]
		S[a]++
		S[b+1]--
	}

	cumsum := []int{0}
	for i := 0; i < len(S); i++ {
		cumsum = append(cumsum, cumsum[i]+S[i])
	}

	var max int
	for _, s := range cumsum {
		if max < s {
			max = s
		}
	}

	return max
}

func TestAnswerABC014Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		AB   [][]int
		want int
	}{
		{"入力例1", 4,
			[][]int{
				{0, 2},
				{2, 3},
				{2, 4},
				{5, 6},
			},
			3,
		},
		{"入力例2", 4,
			[][]int{
				{1000000, 1000000},
				{1000000, 1000000},
				{0, 1000000},
				{1, 1000000},
			},
			4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC014Cその1(tt.N, tt.AB)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
