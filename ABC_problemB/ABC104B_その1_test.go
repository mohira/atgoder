package ABC_problemB

import (
	"atgoder/lib"
	"strings"
	"testing"
)

// [ABC104B - AcCepted](https://atcoder.jp/contests/abc104/tasks/abc104_b)
func AnswerABC104Bその1(S string) string {
	// S の先頭の文字は大文字の A である。
	cond1 := S[0] == 'A'

	// S の先頭から 3 文字目と末尾から 2 文字目の間（両端含む）に大文字の C がちょうど 1 個含まれる。
	cond2 := strings.Count(S[2:len(S)-1], "C") == 1

	// 以上の A, C を除く S のすべての文字は小文字である。
	cond3 := true
	for _, c := range S {
		if c == 'A' || c == 'C' {
			continue
		}
		if !lib.IsLower(c) {
			cond3 = false
		}
	}

	if cond1 && cond2 && cond3 {
		return "AC"
	} else {
		return "WA"
	}
}

func TestAnswerABC104Bその1(t *testing.T) {
	tests := []struct {
		name string
		S    string
		want string
	}{
		{"入力例1", "AtCoder", "AC"},
		{"入力例2", "ACoder", "WA"},
		{"入力例3", "AcycliC", "WA"},
		{"入力例4", "AtCoCo", "WA"},
		{"入力例5", "Atcoder", "WA"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC104Bその1(tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
