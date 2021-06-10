package ABC_problemB

import (
	"testing"
)

// [ABC099B - Stone Monument](https://atcoder.jp/contests/abc099/tasks/abc099_b)
func AnswerABC099Bその1(a, b int) int {
	// 累積和を使って、東西の塔の高さをすべて求めちゃう作戦
	cumsum := make([]int, 999+1)

	for i := 1; i < len(cumsum); i++ {
		cumsum[i] = cumsum[i-1] + i
	}

	for i := 0; i < len(cumsum)-1; i++ {
		A := cumsum[i]
		B := cumsum[i+1]
		if (B - A) == (b - a) {
			return A - a
		}
	}

	return 0
}

func TestAnswerABC099Bその1(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"入力例1", 8, 13, 2},
		{"入力例2", 54, 65, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC099Bその1(tt.a, tt.b)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
