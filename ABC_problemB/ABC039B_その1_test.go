package ABC_problemB

import (
	"math"
	"testing"
)

// [ABC039B - エージェント高橋君](https://atcoder.jp/contests/abc039/tasks/abc039_b)
func AnswerABC039Bその1(X int) int {
	return int(math.Sqrt(math.Sqrt(float64(X))))
}

func TestAnswerABC039Bその1(t *testing.T) {
	tests := []struct {
		name string
		X    int
		want int
	}{
		{"入力例1", 1, 1},
		{"入力例2", 981506241, 177},
		{"入力例3", 390625, 25},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC039Bその1(tt.X)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
