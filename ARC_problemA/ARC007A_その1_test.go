package ARC_problemA

import (
	"strings"
	"testing"
)

// [ARC007A - 帰ってきた器物損壊！高橋君](https://atcoder.jp/contests/arc007/tasks/arc007_1)
func AnswerARC007Aその1(X, s string) string {
	var out strings.Builder
	out.Grow(len(s))

	for _, c := range s {
		if string(c) != X {
			out.WriteRune(c)
		}
	}

	return out.String()
}

func TestAnswerARC007Aその1(t *testing.T) {
	tests := []struct {
		name string
		X    string
		s    string
		want string
	}{
		{"入力例1", "a", "abcdefgabcdefg", "bcdefgbcdefg"},
		{"入力例2", "g", "aassddffgg", "aassddff"},
		{"入力例3", "a", "aaaaa", ""},
		{"入力例4", "l", "qwertyuiopasdfghjklzxcvbnm", "qwertyuiopasdfghjkzxcvbnm"},
		{"入力例5", "d", "qwsdtgcszddddsdfgvbbnj", "qwstgcszsfgvbbnj"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerARC007Aその1(tt.X, tt.s)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
