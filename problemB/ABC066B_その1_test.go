package problemB

import (
	"testing"
)

// [ABC066B - Theater](https://atcoder.jp/contests/abc077/tasks/abc077_b)
func AnswerABC066Bその1(S string) int {
	for i := 1; i < len(S); i++ {
		substr := S[:len(S)-i]

		if is偶文字列(substr) {
			return len(substr)
		}
	}
	return 0
}

func is偶文字列(s string) bool {
	l := len(s)

	left := s[:l/2]
	right := s[l/2:]

	return left == right
}

func TestAnswerABC066Bその1(t *testing.T) {
	tests := []struct {
		name string
		S    string
		want int
	}{
		{"入力例1", "abaababaab", 6},
		{"入力例2", "xxxx", 2},
		{"入力例3", "abcabcabcabc", 6},
		{"入力例4", "akasakaakasakasakaakas", 14},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC066Bその1(tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
