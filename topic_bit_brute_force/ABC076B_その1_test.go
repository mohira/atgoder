package topic_bit_brute_force

import (
	"math"
	"testing"
)

// [ABC076B - Addition and Multiplication](https://atcoder.jp/contests/abc076/tasks/abc076_b)
func AnswerABC076Bその1(N, K int) int {
	// ビット全探索
	var ans = math.MaxInt64

	for bit := 0; bit < (1 << N); bit++ {
		tmp := 1
		for i := 0; i < N; i++ {
			if (bit>>i)&1 == 1 {
				// 操作A
				tmp *= 2
			} else {
				// 操作B
				tmp += K
			}
		}
		ans = min(ans, tmp)
	}

	return ans
}

func TestAnswerABC076Bその1(t *testing.T) {
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
			got := AnswerABC076Bその1(tt.N, tt.K)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
