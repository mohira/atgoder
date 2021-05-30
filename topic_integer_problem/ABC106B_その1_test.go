package topic_integer_problem

import (
	"testing"
)

// [ABC106B - 105](https://atcoder.jp/contests/abc106/tasks/abc106_b)
func AnswerABC106Bその1(N int) int {
	var ans int

	for odd := 1; odd <= N; odd += 2 {
		divisorsCount := 0

		for j := 1; j*j <= odd; j++ {
			if odd%j == 0 {
				divisorsCount++
				if odd/j != j {
					divisorsCount++
				}
			}
		}

		if divisorsCount == 8 {
			ans++
		}
	}

	return ans
}

func TestAnswerABC106Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{"入力例1", 105, 1},
		{"入力例2", 7, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC106Bその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
