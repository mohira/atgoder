package similar

import (
	"strings"
	"testing"
)

// [ABC085A - Already 2018](https://atcoder.jp/contests/abc085/tasks/abc085_a)
func AnswerABC085A(S string) string {
	return strings.Replace(S, "2017", "2018", 1)
}

func TestABC085A(t *testing.T) {
	tests := []struct {
		name string
		S    string
		want string
	}{
		{"入力例1", "2017/01/07", "2018/01/07"},
		{"入力例2", "2017/01/31", "2018/01/31"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC085A(tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
