package ABC_problemA

import (
	"testing"
)

// [ABC047A  - キャンディーと2人の子供](https://atcoder.jp/contests/abc047/tasks/abc047_a)
func AnswerABC047Aその1(a, b, c int) string {
	if a == b+c || b == c+a || c == a+b {
		return "Yes"
	} else {
		return "No"
	}
}

func TestAnswerABC047Aその1(t *testing.T) {
	tests := []struct {
		name    string
		a, b, c int
		want    string
	}{
		{"入力例1", 10, 30, 20, "Yes"},
		{"入力例2", 30, 30, 100, "No"},
		{"入力例3", 56, 25, 31, "Yes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC047Aその1(tt.a, tt.b, tt.c)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
