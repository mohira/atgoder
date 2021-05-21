package problemB

import (
	"strings"
	"testing"
)

// [ABC154B - I miss you...](https://atcoder.jp/contests/abc154/tasks/abc154_b)
func AnswerABC154Bその1(S string) string {
	return strings.Repeat("x", len(S))
}

func TestAnswerABC154Bその1(t *testing.T) {
	tests := []struct {
		name string
		S    string
		want string
	}{
		{"入力例1", "sardine", "xxxxxxx"},
		{"入力例2", "xxxx", "xxxx"},
		{"入力例3", "gone", "xxxx"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC154Bその1(tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
