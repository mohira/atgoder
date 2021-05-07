package similar

import (
	"testing"
)

// [ABC064A - RGB Cards](https://atcoder.jp/contests/abc064/tasks/abc064_a)
func AnswerABC064Aその1(r int, g int, b int) string {
	n := 100*r + 10*g + b

	if n%4 == 0 {
		return "YES"
	} else {
		return "NO"
	}
}

func TestAnswerABC064Aその1(t *testing.T) {
	tests := []struct {
		name string
		r    int
		g    int
		b    int
		want string
	}{
		{"入力例1", 4, 3, 2, "YES"},
		{"入力例2", 2, 3, 4, "NO"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC064Aその1(tt.r, tt.g, tt.b)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
