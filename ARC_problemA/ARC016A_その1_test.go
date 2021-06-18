package ARC_problemA

import (
	"testing"
)

// [ARC016A - クイズゲーム](https://atcoder.jp/contests/arc016/tasks/arc016_1)
func AnswerARC016Aその1(N, M int) int {
	if M == 1 {
		return 2
	}
	return 1
}

func TestAnswerARC016Aその1(t *testing.T) {
	tests := []struct {
		name string
		N, M int
		want int
	}{
		{"入力例1", 4, 4, 1},
		{"入力例2", 2, 1, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerARC016Aその1(tt.N, tt.M)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
