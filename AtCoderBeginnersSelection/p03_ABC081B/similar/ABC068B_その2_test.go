package similar

import (
	"testing"
)

// [ABC068B - Break Number](https://atcoder.jp/contests/abc068/tasks/abc068_b)
func AnswerABC068Bその2(N int) int {
	// そもそも解の候補は、100以下の2の累乗にしぼられる [1, 2, 4, 8, 16, 32, 64]
	// 候補のうち、N以下 かつ 最大 を求めればOK
	var answerCandidates = []int{64, 32, 16, 8, 4, 2, 1}

	var ans = -1

	for _, candidate := range answerCandidates {
		if candidate <= N {
			return candidate
		}
	}

	return ans
}

func TestAnswerABC068Bその2(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{"入力例1", 7, 4},
		{"入力例2", 32, 32},
		{"入力例3", 1, 1},
		{"入力例3", 99, 64},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC068Bその2(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}

		})

	}
}
