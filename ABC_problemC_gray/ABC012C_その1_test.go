package ABC_problemC_gray

import (
	"fmt"
	"strings"
	"testing"
)

// [ABC012C - 九九足し算](https://atcoder.jp/contests/abc012/tasks/abc012_3)
func AnswerABC012Cその1(N int) string {
	const total = 2025
	x := total - N

	var out []string

	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if i*j == x {
				out = append(out, fmt.Sprintf("%d x %d", i, j))
			}
		}
	}
	return strings.Join(out, "\n")
}

func TestAnswerABC012Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want string
	}{
		{"入力例1", 2013, "2 x 6\n3 x 4\n4 x 3\n6 x 2"},
		{"入力例2", 2024, "1 x 1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC012Cその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
