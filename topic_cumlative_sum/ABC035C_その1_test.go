package topic_cumlative_sum

import (
	"testing"
)

// [ABC035C - オセロ](https://atcoder.jp/contests/tokiomarine2020/tasks/tokiomarine2020_c)
func AnswerABC035Cその1(N, Q int, LR [][]int) string {
	var S = make([]int, N+1)

	for _, lr := range LR {
		l, r := lr[0], lr[1]
		l--
		r--
		S[l]++
		S[r+1]--
	}

	cumsum := make([]int, 0, len(S)+1)
	cumsum = append(cumsum, 0)
	for i := 0; i < len(S); i++ {
		cumsum = append(cumsum, cumsum[i]+S[i])
	}

	var ans string
	for i := 1; i < len(S); i++ {
		flipCount := cumsum[i]
		if flipCount%2 == 0 {
			ans += "0"
		} else {
			ans += "1"
		}
	}

	return ans
}

func TestAnswerABC035Cその1(t *testing.T) {
	tests := []struct {
		name string
		N, Q int
		LR   [][]int
		want string
	}{
		{"入力例1",
			5, 4,
			[][]int{
				{5, 4},
				{1, 4},
				{2, 5},
				{3, 3},
				{1, 5},
			},
			"01010"},
		{"入力例2",
			20, 8,
			[][]int{
				{1, 8},
				{4, 13},
				{8, 8},
				{3, 18},
				{5, 20},
				{19, 20},
				{2, 7},
				{4, 9},
			},
			"10110000011110000000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC035Cその1(tt.N, tt.Q, tt.LR)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
