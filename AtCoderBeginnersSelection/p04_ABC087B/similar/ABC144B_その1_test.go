package similar

import (
	"testing"
)

// [ABC144B - 81](https://atcoder.jp/contests/abc144/tasks/abc144_b)
func AnswerABC144Bその1(N int) string {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if i*j == N {
				return "Yes"
			}
		}
	}

	return "No"
}

func TestAnswerABC144Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want string
	}{
		{"入力例1", 10, "Yes"},
		{"入力例2", 50, "No"},
		{"入力例3", 81, "Yes"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC144Bその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}

		})

	}
}
