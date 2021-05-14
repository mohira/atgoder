package problemB

import (
	"testing"
)

// [ABC184B - Quizzes](https://atcoder.jp/contests/abc184/tasks/abc184_b)
func AnswerABC184Bその1(N, X int, S string) int {
	point := X

	for _, c := range S {
		if c == 'x' {
			point = max(point-1, 0)
		} else {
			point++
		}
	}

	return point
}

func TestAnswerABC184Bその1(t *testing.T) {
	tests := []struct {
		name string
		N, X int
		S    string
		want int
	}{
		{"入力例1", 3, 0, "xox", 0},
		{"入力例2", 20, 199999, "oooooooooxoooooooooo", 200017},
		{"入力例2", 20, 10, "xxxxxxxxxxxxxxxxxxxx", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC184Bその1(tt.N, tt.X, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
