package ABC_problemB

import (
	"strings"
	"testing"
)

// [ABC196B - Round Down](https://atcoder.jp/contests/abc196/tasks/abc196_b)
func AnswerABC196Bその1(X string) string {
	return strings.Split(X, ".")[0]
}

func TestAnswerABC196Bその1(t *testing.T) {
	tests := []struct {
		name string
		X    string
		want string
	}{
		{"入力例1", "5.90", "5"},
		{"入力例2", "0", "0"},
		{"入力例3", "84939825309432908832902189.9092309409809091329", "84939825309432908832902189"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC196Bその1(tt.X)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
