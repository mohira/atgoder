package ARC_problemA

import (
	"strings"
	"testing"
)

// [ARC045A - スペース高橋君](https://atcoder.jp/contests/arc045/tasks/arc045_a)
func AnswerARC045Aその1(S string) string {
	words := strings.Split(S, " ")

	out := make([]string, len(words))
	for i, word := range words {
		out[i] = takahashiTask(word)
	}
	return strings.Join(out, " ")
}

func takahashiTask(word string) string {
	switch word {
	case "Left":
		return "<"
	case "Right":
		return ">"
	case "AtCoder":
		return "A"
	default:
		panic("おかしいよ")
	}
}

func TestAnswerARC045Aその1(t *testing.T) {
	tests := []struct {
		name string
		S    string
		want string
	}{
		{"入力例1", "Left Right AtCoder", "< > A"},
		{"入力例2", "Left Left Right Right AtCoder", "< < > > A"},
		{"入力例3", "Right Right AtCoder Left Left AtCoder", "> > A < < A"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerARC045Aその1(tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
