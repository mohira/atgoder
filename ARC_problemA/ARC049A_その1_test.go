package ARC_problemA

import (
	"strings"
	"testing"
)

// [ARC049A - "強調"](https://atcoder.jp/contests/arc049/tasks/arc049_1)
func AnswerARC049Aその1(S string, A, B, C, D int) string {
	var out strings.Builder
	out.Grow(len(S) + 4)

	for i, c := range S {
		if i == A || i == B || i == C || i == D {
			out.WriteRune('"')
		}
		out.WriteRune(c)
	}

	if len(S) == D {
		out.WriteRune('"')
	}

	return out.String()
}

func TestAnswerARC049Aその1(t *testing.T) {
	tests := []struct {
		name       string
		S          string
		A, B, C, D int
		want       string
	}{
		{
			"入力例1",
			"MinnnahaNakayoshi",
			0, 6, 8, 17,
			"\"Minnna\"ha\"Nakayoshi\"",
		},
		{
			"入力例2",
			"Niwawo_Kakemeguru_Chokudai",
			11, 17, 18, 26,
			"Niwawo_Kake\"meguru\"_\"Chokudai\"",
		},
		{
			"入力例3",
			"___",
			0, 1, 2, 3,
			"\"_\"_\"_\"",
		},
		{
			"",
			"abcde",
			0, 1, 2, 3,
			"\"a\"b\"c\"de",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerARC049Aその1(tt.S, tt.A, tt.B, tt.C, tt.D)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
