package ABC_problemB

import (
	"testing"
)

// [ABC063B - Theater](https://atcoder.jp/contests/abc077/tasks/abc077_b)
func AnswerABC063Bその1(S string) string {
	bucket := make(map[rune]int)
	for _, c := range S {
		if _, ok := bucket[c]; ok {
			return "no"
		} else {
			bucket[c] = 1
		}
	}
	return "yes"
}

func TestAnswerABC063Bその1(t *testing.T) {
	tests := []struct {
		name string
		S    string
		want string
	}{
		{"入力例1", "uncopyrightable", "yes"},
		{"入力例2", "different", "no"},
		{"入力例3", "no", "yes"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC063Bその1(tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
