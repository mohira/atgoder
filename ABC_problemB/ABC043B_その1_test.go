package ABC_problemB

import (
	"testing"
)

// [ABC043B - バイナリハックイージー](https://atcoder.jp/contests/abc043/tasks/abc043_b)
func AnswerABC043Bその1(S string) string {
	var ans string

	for _, c := range S {
		if c == '0' {
			ans += "0"
		}
		if c == '1' {
			ans += "1"
		}

		if c == 'B' && len(ans) >= 1 {
			ans = ans[:len(ans)-1]
		}
	}

	return ans
}

func TestAnswerABC043Bその1(t *testing.T) {
	tests := []struct {
		name string
		S    string
		want string
	}{
		{"入力例1", "01B0", "00"},
		{"入力例2", "0BB1", "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC043Bその1(tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
