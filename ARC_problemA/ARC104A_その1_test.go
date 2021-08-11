package ARC_problemA

import (
	"fmt"
	"testing"
)

// [ARC104A - Plus Minus](https://atcoder.jp/contests/arc104/tasks/arc104_1)
func AnswerARC104Aその1(A, B int) string {
	for x := -100; x <= 100; x++ {
		for y := -100; y <= 100; y++ {
			if x+y == A && x-y == B {
				return fmt.Sprintf("%d %d", x, y)
			}
		}
	}
	return ""
}

func TestAnswerARC104Aその1(t *testing.T) {
	tests := []struct {
		name string
		A, B int
		want string
	}{
		{"入力例1", 2, -2, "0 2"},
		{"入力例2", 3, 1, "2 1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerARC104Aその1(tt.A, tt.B)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
