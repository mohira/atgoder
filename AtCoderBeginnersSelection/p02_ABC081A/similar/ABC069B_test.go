package similar

import (
	"fmt"
	"testing"
)

// [ABC069B - i18n](https://atcoder.jp/contests/abc069/tasks/abc069_b)
func AnswerABC069B(s string) string {
	head := string(s[0])
	tail := string(s[len(s)-1])

	length := len(s) - 2

	return fmt.Sprintf("%s%d%s", head, length, tail)
}

func TestABC069B(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"入力例1", "internationalization", "i18n"},
		{"入力例2", "smiles", "s4s"},
		{"入力例3", "xyz", "x1z"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC069B(tt.s)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
