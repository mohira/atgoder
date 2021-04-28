package similar

import (
	"testing"
)

// [ABC105B - Cakes and Donuts](https://atcoder.jp/contests/abc105/tasks/abc105_b)
func AnswerABC105Bその1(N int) string {
	var cakePrice = 4
	var donutPrice = 7

	for i := 0; i <= N/cakePrice; i++ {
		for j := 0; j <= N/donutPrice; j++ {
			totalPrice := cakePrice*i + donutPrice*j

			if totalPrice == N {
				return "Yes"
			}
		}
	}

	return "No"
}

func TestAnswerABC105Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want string
	}{
		{"入力例1", 11, "Yes"},
		{"入力例2", 40, "Yes"},
		{"入力例3", 3, "No"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC105Bその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}

		})

	}
}
