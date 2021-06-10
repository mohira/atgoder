package ABC_problemB

import (
	"strings"
	"testing"
)

// [ABC084B - Postal Code](https://atcoder.jp/contests/abc084/tasks/abc084_b)
func AnswerABC084Bその1(A, B int, S string) string {
	// ハイフンがただ1つだけ含まれているかどうか
	if strings.Count(S, "-") != 1 {
		return "No"
	}

	// ハイフンの位置が正しいかどうか
	idx := strings.Index(S, "-")
	if idx+1 != A+1 {
		return "No"
	}

	// ハイフン以外の文字が数字で構成されているかどうか
	for _, s := range strings.Split(S, "-") {
		for _, c := range s {
			n := c - '0'
			if !(0 <= n && n <= 9) {
				return "No"
			}
		}
	}

	return "Yes"
}

func TestAnswerABC084Bその1(t *testing.T) {
	tests := []struct {
		name string
		A, B int
		S    string
		want string
	}{
		{"入力例1", 3, 4, "269-6650", "Yes"},
		{"入力例2", 1, 1, "---", "No"},
		{"入力例3", 1, 2, "7444", "No"},
		{"ハイフンが複数", 3, 4, "269-6-650", "No"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC084Bその1(tt.A, tt.B, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
