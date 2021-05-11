package topic_bit_brute_force

import (
	"testing"
)

// [ABC018B - 文字列の反転](https://atcoder.jp/contests/abc018/tasks/abc018_2)
func AnswerABC018Bその1(S string, N int, LR [][]int) string {
	for i := 0; i < N; i++ {
		l, r := LR[i][0], LR[i][1]
		l--
		r--

		tmp := ""
		for a := 0; a < l; a++ {
			tmp += string(S[a])
		}

		// 反転部分
		for b := r; b >= l; b-- {
			tmp += string(S[b])
		}

		for c := r + 1; c < len(S); c++ {
			tmp += string(S[c])
		}

		S = tmp
	}

	return S
}

func TestAnswerABC018Bその1(t *testing.T) {
	tests := []struct {
		name string
		S    string
		N    int
		LR   [][]int
		want string
	}{
		{"入力例1",
			"abcdef",

			2,
			[][]int{
				{3, 5},
				{1, 4},
			},
			"debacf",
		},

		{"入力例2",
			"redcoat",
			3,
			[][]int{
				{1, 7},
				{1, 2},
				{3, 4},
			},
			"atcoder",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC018Bその1(tt.S, tt.N, tt.LR)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
