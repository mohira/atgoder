package ABC_problemB

import (
	"atgoder/lib"
	"testing"
)

// [ABC013B - 錠](https://atcoder.jp/contests/abc013/tasks/abc013_2)
func AnswerABC013Bその1(a, b int) int {
	return lib.Min(lib.AbsInt(a-b), lib.Min(b+10-a, a+10-b))
}

func TestAnswerABC013Bその1(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"入力例1", 4, 6, 2},
		{"入力例2", 6, 4, 2},
		{"入力例3", 8, 1, 3},
		{"入力例4", 1, 8, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC013Bその1(tt.a, tt.b)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
