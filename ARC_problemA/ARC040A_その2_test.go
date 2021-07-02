package ARC_problemA

import (
	"strings"
	"testing"
)

// [ARC040A - 床塗り](https://atcoder.jp/contests/arc040/tasks/arc040_a)
func AnswerARC040Aその2(N int, S []string) string {
	var concatenated string

	for _, s := range S {
		concatenated += s
	}

	red := strings.Count(concatenated, "R")
	blue := strings.Count(concatenated, "B")

	if blue < red {
		return "TAKAHASHI"
	}

	if red < blue {
		return "AOKI"
	}

	return "DRAW"
}

func TestAnswerARC040Aその2(t *testing.T) {
	tests := []struct {
		name string
		N    int
		S    []string
		want string
	}{
		{
			"入力例1",
			4,
			[]string{
				"R.RB",
				"RR.B",
				"BRBB",
				"RRB.",
			},
			"TAKAHASHI",
		},
		{
			"入力例2",
			2,
			[]string{
				"..",
				"..",
			},
			"DRAW",
		},
		{
			"入力例3",
			3,
			[]string{
				"BRB",
				"RBR",
				"BRB",
			},
			"AOKI",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerARC040Aその2(tt.N, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
