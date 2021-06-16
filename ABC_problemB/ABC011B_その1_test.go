package ABC_problemB

import (
	"strings"
	"testing"
)

// [ABC011B - 名前の確認](https://atcoder.jp/contests/abc011/tasks/abc011_2)
func AnswerABC011Bその1(S string) string {
	var ans string

	ans += strings.ToUpper(string(S[0]))
	ans += strings.ToLower(S[1:])

	return ans
}

func TestAnswerABC011Bその1(t *testing.T) {
	tests := []struct {
		name string
		S    string
		want string
	}{
		{"入力例1", "taKahAshI", "Takahashi"},
		{"入力例2", "A", "A"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC011Bその1(tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
