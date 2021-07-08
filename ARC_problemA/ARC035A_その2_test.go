package ARC_problemA

import (
	"atgoder/lib"
	"testing"
)

// [ARC035A - 高橋くんと回文](https://atcoder.jp/contests/arc035/tasks/arc035_a)
func AnswerARC035Aその2(S string) string {
	reverseStr := lib.ReverseStr(S)
	ok := true

	for i := 0; i < len(S); i++ {
		head := S[i]
		tail := reverseStr[i]

		if head == tail || head == '*' || tail == '*' {
			continue
		}

		ok = false

	}

	if ok {
		return "YES"
	} else {
		return "NO"
	}
}

func TestAnswerARC035Aその2(t *testing.T) {
	tests := []struct {
		name string
		S    string
		want string
	}{
		{"入力例1", "ab*", "YES"},
		{"入力例2", "abc", "NO"},
		{"入力例3", "a*bc*", "YES"},
		{"入力例4", "***", "YES"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerARC035Aその2(tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
