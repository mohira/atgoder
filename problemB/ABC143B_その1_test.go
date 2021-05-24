package problemB

import (
	"testing"
)

// [ABC143B - TAKOYAKI FESTIVAL 2019](https://atcoder.jp/contests/abc143/tasks/abc143_b)
func AnswerABC143Bその1(N int, D []int) int {
	var ans int

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			ans += D[i] * D[j]
		}
	}

	return ans
}

func TestAnswerABC143Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		D    []int
		want int
	}{
		{"入力例1", 3, []int{3, 1, 2}, 11},
		{"入力例2", 7, []int{5, 0, 7, 8, 3, 3, 2}, 312},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC143Bその1(tt.N, tt.D)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
