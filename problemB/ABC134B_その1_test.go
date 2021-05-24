package problemB

import (
	"testing"
)

// [ABC134B - Golden Apple](https://atcoder.jp/contests/abc134/tasks/abc134_b)
func AnswerABC134Bその1(N, D int) int {
	var count int
	for i := 0; i < N; i++ {
		i += D + D // 最大守備範囲
		count++
	}
	return count
}

func TestAnswerABC134Bその1(t *testing.T) {
	tests := []struct {
		name string
		N, D int
		want int
	}{
		{"入力例1", 6, 2, 2},
		{"入力例2", 14, 3, 2},
		{"入力例3", 20, 4, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC134Bその1(tt.N, tt.D)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
