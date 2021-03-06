package p02_ABC081A

import (
	"strings"
	"testing"
)

// [ABC081A - Placing Marbles](https://atcoder.jp/contests/abc081/tasks/abc081_a)
func AnswerABC081Aその1(s1s2s3 string) int {
	return strings.Count(s1s2s3, "1")
}

func TestAnswerABC081Aその1(t *testing.T) {
	tests := []struct {
		name   string
		s1s2s3 string
		want   int
	}{
		{"入力例1", "101", 2},
		{"入力例2", "000", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC081Aその1(tt.s1s2s3)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
