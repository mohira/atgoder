package ARC_problemA

import (
	"testing"
)

// [ARC052A - 何期生？](https://atcoder.jp/contests/arc052/tasks/arc052_1)
func AnswerARC052Aその1(S string) string {
	var ans string

	for _, c := range S {
		n := int(c - '0')
		if 0 <= n && n <= 9 {
			ans += string(c)
		}
	}

	return ans
}

func TestAnswerARC052Aその1(t *testing.T) {
	tests := []struct {
		name string
		S    string
		want string
	}{
		{"入力例1", "chokudai55", "55"},
		{"入力例2", "cho9dai", "9"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerARC052Aその1(tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
