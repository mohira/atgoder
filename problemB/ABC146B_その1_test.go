package problemB

import (
	"testing"
)

// [ABC146B - ROT N](https://atcoder.jp/contests/abc146/tasks/abc146_b)
func AnswerABC146Bその1(N int, S string) string {
	var ans string

	for _, c := range S {
		rot := c + rune(N%26)

		if 'A' <= rot && rot <= 'Z' {
			ans += string(rot)
		} else {
			ans += string(rot - 26)
		}
	}

	return ans
}

func TestAnswerABC146Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		S    string
		want string
	}{
		{"入力例1", 2, "ABCXYZ", "CDEZAB"},
		{"入力例2", 0, "ABCXYZ", "ABCXYZ"},
		{"入力例3", 13, "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "NOPQRSTUVWXYZABCDEFGHIJKLM"},

		{"オーバーするやつ (0<=N<=26)だけど気になったので追加", 2 + 26, "ABCXYZ", "CDEZAB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC146Bその1(tt.N, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
