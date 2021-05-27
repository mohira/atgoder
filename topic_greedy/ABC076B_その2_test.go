package topic_greedy

import (
	"testing"
)

// [ABC076B - Addition and Multiplication](https://atcoder.jp/contests/abc076/tasks/abc076_b)
func AnswerABC076Bその2(N, K int) int {
	// 貪欲法
	var ans = 1

	for i := 0; i < N; i++ {
		ans = min(ans+K, ans*2)
	}

	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func TestAnswerABC076Bその2(t *testing.T) {
	tests := []struct {
		name string
		N, K int
		want int
	}{
		{"入力例1", 4, 3, 10},
		{"入力例2", 10, 10, 76},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC076Bその2(tt.N, tt.K)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
