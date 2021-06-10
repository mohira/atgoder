package ABC_problemB

import (
	"testing"
)

// [ABC141B - Tap Dance](https://atcoder.jp/contests/abc141/tasks/abc141_b)
func AnswerABC141Bその1(S string) string {
	for i, c := range S {
		if i%2 == 0 {
			if c == 'L' {
				return "No"
			}
		} else {
			if c == 'R' {
				return "No"
			}
		}
	}
	return "Yes"
}

func TestAnswerABC141Bその1(t *testing.T) {
	tests := []struct {
		name string
		S    string
		want string
	}{
		{"入力例1", "RUDLUDR", "Yes"},
		{"入力例2", "DULL", "No"},
		{"入力例3", "UUUUUUUUUUUUUUU", "Yes"},
		{"入力例4", "ULURU", "No"},
		{"入力例5", "RDULULDURURLRDULRLR", "Yes"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC141Bその1(tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
