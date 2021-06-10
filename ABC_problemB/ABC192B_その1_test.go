package ABC_problemB

import (
	"atgoder/lib"
	"testing"
)

// [ABC192B - uNrEaDaBlE sTrInG](https://atcoder.jp/contests/abc192/tasks/abc192_b)
func AnswerABC192Bその1(S string) string {
	for i, c := range S {
		if i%2 == 0 && !lib.IsLower(c) {
			return "No"
		}

		if i%2 == 1 && !lib.IsUpper(c) {
			return "No"
		}
	}

	return "Yes"
}

func TestAnswerABC192Bその1(t *testing.T) {
	tests := []struct {
		name string
		S    string
		want string
	}{
		{"入力例1", "dIfFiCuLt", "Yes"},
		{"入力例2", "eASY", "No"},
		{"入力例3", "a", "Yes"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC192Bその1(tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
