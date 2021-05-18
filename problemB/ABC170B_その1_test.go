package problemB

import (
	"testing"
)

// [ABC170B - Crane and Turtle](https://atcoder.jp/contests/abc170/tasks/abc170_b)
func AnswerABC170Bその1(X, Y int) string {
	for turtle := 0; turtle <= 25; turtle++ {
		for crane := 0; crane <= 50; crane++ {
			totalAnimals := turtle + crane
			totalLegs := 4*turtle + 2*crane

			if totalAnimals == X && totalLegs == Y {
				return "Yes"
			}
		}
	}
	return "No"
}

func TestAnswerABC170Bその1(t *testing.T) {
	tests := []struct {
		name string
		X, Y int
		want string
	}{
		{"入力例1", 3, 8, "Yes"},
		{"入力例2", 2, 100, "No"},
		{"入力例3", 1, 2, "Yes"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC170Bその1(tt.X, tt.Y)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
