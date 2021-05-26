package problemB

import (
	"testing"
)

// [ABC099B - Stone Monument](https://atcoder.jp/contests/abc099/tasks/abc099_b)
func AnswerABC099Bその2(a, b int) int {
	return (b-a)*(b-a+1)/2 - b
}

func TestAnswerABC099Bその2(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"入力例1", 8, 13, 2},
		{"入力例2", 54, 65, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC099Bその2(tt.a, tt.b)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
