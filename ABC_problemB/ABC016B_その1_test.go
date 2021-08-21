package ABC_problemB

import (
	"testing"
)

// [ABC016B - A±B Problem](https://atcoder.jp/contests/abc016/tasks/abc016_2)
func AnswerABC016Bその1(a, b, c int) string {
	if a+b == c && a-b == c {
		return "?"
	}
	if a+b == c {
		return "+"
	}
	if a-b == c {
		return "-"
	}

	return "!"
}

func TestAnswerABC016Bその1(t *testing.T) {
	tests := []struct {
		name    string
		a, b, c int
		want    string
	}{
		{"入力例1", 1, 0, 1, "?"},
		{"入力例2", 1, 1, 2, "+"},
		{"入力例3", 1, 1, 0, "-"},
		{"入力例4", 1, 1, 1, "!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC016Bその1(tt.a, tt.b, tt.c)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
