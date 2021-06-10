package ABC_problemB

import (
	"strings"
	"testing"
)

// [ABC058B - Comparison](https://atcoder.jp/contests/abc058/tasks/abc058_b)
func AnswerABC058Bその1(O, E string) string {
	lenO := len(O)
	lenE := len(E)

	var out strings.Builder
	out.Grow(lenO + lenE)

	for i := 0; i < lenE; i++ {
		out.WriteByte(O[i])
		out.WriteByte(E[i])
	}

	if lenO-lenE == 1 {
		out.WriteByte(O[lenO-1])
	}

	return out.String()
}

func TestAnswerABC058Bその1(t *testing.T) {
	tests := []struct {
		name string
		O, E string
		want string
	}{
		{"入力例1: 文字数が等しい", "xyz", "abc", "xaybzc"},
		{"入力例2: 文字数が等しくない", "atcoderbeginnercontest", "atcoderregularcontest", "aattccooddeerrbreeggiunlnaerrccoonntteesstt"},
		{"", "ac", "b", "abc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC058Bその1(tt.O, tt.E)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
