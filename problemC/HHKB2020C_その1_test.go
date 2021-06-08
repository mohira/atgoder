package problemC

import (
	"testing"
)

// [HHKB2020C - Neq Min](https://atcoder.jp/contests/hhkb2020/tasks/hhkb2020_c)
func AnswerHHKB2020Cその1(N int, P []int) string {
	return ""
}

func TestAnswerHHKB2020Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		P    []int
		want string
	}{
		{"入力例1", 4, []int{1, 1, 0, 2}, "0\n0\n2\n3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerHHKB2020Cその1(tt.N, tt.P)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
