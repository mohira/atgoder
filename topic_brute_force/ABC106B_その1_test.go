package topic_brute_force

import (
	"testing"
)

// [ABC106B - 105](https://atcoder.jp/contests/abc106/tasks/abc106_b)
func AnswerABC106Bその1(N int) int {
	var ans int
	for i := 1; i <= N; i += 2 {
		if countDivisors(i) == 8 {
			ans++
		}
	}
	return ans
}

func countDivisors(n int) int {
	var count int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	return count
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
