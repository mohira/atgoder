package topic_amari

import (
	"atgoder/lib"
	"math"
	"testing"
)

// [ABC156C - Rally](https://atcoder.jp/contests/abc156/tasks/abc156_c)
func AnswerABC156Cその1(N int, X []int) int {
	var minPoint = math.MaxInt64

	// 1<=xi<=100 なので、1から調べればいい
	for i := 1; i <= 100; i++ {
		point := 0
		for _, x := range X {
			point += (x - i) * (x - i)
		}

		minPoint = lib.Min(minPoint, point)
	}

	return minPoint
}

func TestAnswerABC156Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		X    []int
		want int
	}{
		{"入力例1", 2, []int{1, 4}, 5},
		{"入力例2", 7, []int{14, 14, 2, 13, 56, 2, 37}, 2354},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC156Cその1(tt.N, tt.X)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
