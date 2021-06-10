package ABC_problemB

import (
	"testing"
)

// [ABC147B - Palindrome-philsaia](https://atcoder.jp/contests/abc147/tasks/abc147_b)
func AnswerABC147Bその1(S string) int {
	var count int

	for i := 0; i < len(S)/2; i++ {
		head := i
		tail := len(S) - 1 - i

		if S[head] != S[tail] {
			count++
		}
	}

	return count
}

func TestAnswerABC147Bその1(t *testing.T) {
	tests := []struct {
		name string
		S    string
		want int
	}{
		{"入力例1", "redcoder", 1},
		{"入力例2", "vvvvvv", 0},
		{"入力例3: 奇数の場合", "abcdabc", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC147Bその1(tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
