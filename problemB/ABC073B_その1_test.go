package problemB

import (
	"testing"
)

// [ABC073B - Theater](https://atcoder.jp/contests/abc077/tasks/abc077_b)
func AnswerABC073Bその1(N int, LR [][]int) int {
	var ans int
	for i := 0; i < N; i++ {
		l := LR[i][0]
		r := LR[i][1]

		ans += r - l + 1
	}
	return ans
}

func TestAnswerABC073Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		LR   [][]int
		want int
	}{
		{"入力例1",
			1,
			[][]int{
				{24, 30},
			},
			7,
		},
		{"入力例2",
			2,
			[][]int{
				{6, 8},
				{3, 3},
			},
			4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC073Bその1(tt.N, tt.LR)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
