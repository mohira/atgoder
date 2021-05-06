package topic_brute_force

import (
	"testing"
)

// [ABC122B - ATCoder](https://atcoder.jp/contests/abc122/tasks/abc122_b)
func AnswerABC122Bその1(S string) int {
	var maxLength int

	for i := 0; i < len(S); i++ {
		for j := i + 1; j < len(S)+1; j++ {
			substr := S[i:j]

			if isACGT文字列(substr) {
				maxLength = max(maxLength, len(substr))
			}
		}
	}

	return maxLength
}

func max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func isACGT文字列(s string) bool {
	for _, c := range s {
		if !(c == 'A' || c == 'C' || c == 'G' || c == 'T') {
			return false
		}
	}
	return true
}

func TestAnswerABC122Bその1(t *testing.T) {
	tests := []struct {
		name string
		S    string
		want int
	}{
		{"入力例1", "ATCODER", 3},
		{"入力例2", "HATAGAYA", 5},
		{"入力例3", "SHINJUKU", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC122Bその1(tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
