package ARC_problemA

import (
	"testing"
)

// [ARC035A - 高橋くんと回文](https://atcoder.jp/contests/arc035/tasks/arc035_a)
func AnswerARC035Aその1(S string) string {
	N := len(S)
	ok := true

	for i := 0; i < N/2; i++ {
		head, tail := S[i], S[N-i-1]

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

func TestAnswerARC035Aその1(t *testing.T) {
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
			got := AnswerARC035Aその1(tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
