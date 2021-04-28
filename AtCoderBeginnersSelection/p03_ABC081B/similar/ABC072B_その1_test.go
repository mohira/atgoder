package similar

import (
	"testing"
)

// [ABC072B - OddString](https://atcoder.jp/contests/abc072/tasks/abc072_b)
func AnswerABC072Bその1(s string) string {
	var ans string
	for i := 0; i < len(s); i += 2 {
		ans += string(s[i])
	}
	return ans
}

func TestAnswerABC072Bその1(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"入力例1", "atcoder", "acdr"},
		{"入力例2", "aaaa", "aa"},
		{"入力例3", "fukuokayamaguchi", "fkoaaauh"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC072Bその1(tt.s)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}

		})

	}
}
