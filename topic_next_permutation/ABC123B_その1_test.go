package topic_next_permutation

import (
	"testing"
)

// [ABC123B - Five Dishes](https://atcoder.jp/contests/abc123/tasks/abc123_b)
func AnswerABC123Bその2(A, B, C, D, E int) int {
	// 順列全探索でやる
	return 0
}

func TestAnswerABC123Bその2(t *testing.T) {
	tests := []struct {
		name          string
		A, B, C, D, E int
		want          int
	}{
		{"入力例1", 29, 20, 7, 35, 120, 215},
		{"入力例2", 101, 86, 119, 108, 57, 481},
		{"入力例3 最小1桁が複数でるパターン", 5, 10, 5, 0, 0, 25},
		{"入力例4 すべての下1桁が0", 10, 20, 30, 40, 50, 150},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC123Bその2(tt.A, tt.B, tt.C, tt.D, tt.E)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
