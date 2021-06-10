package ABC_problemB

import (
	"atgoder/lib"
	"testing"
)

// [ABC159B - String Palindrome](https://atcoder.jp/contests/abc159/tasks/abc159_b)
func AnswerABC159Bその1(S string) string {
	N := len(S)
	sLeft := S[:(N-1)/2]
	sRight := S[(N+3)/2-1:]

	if lib.IsPalindrome(S) && lib.IsPalindrome(sLeft) && lib.IsPalindrome(sRight) {
		return "Yes"
	} else {
		return "No"
	}
}

func TestAnswerABC159Bその1(t *testing.T) {
	tests := []struct {
		name string
		X    string
		want string
	}{
		{"入力例1", "akasaka", "Yes"},
		{"入力例2", "level", "No"},
		{"入力例3", "atcoder", "No"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC159Bその1(tt.X)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
