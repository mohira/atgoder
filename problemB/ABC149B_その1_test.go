package problemB

import (
	"fmt"
	"testing"
)

// [ABC149B - Achieve the Goal](https://atcoder.jp/contests/abc149/tasks/abc149_b)
func AnswerABC149Bその1(A, B, K int) string {
	A, K = max(0, A-K), max(0, K-A)
	B, K = max(0, B-K), max(0, K-B)

	return fmt.Sprintf("%d %d", A, B)
}

func TestAnswerABC149Bその1(t *testing.T) {
	tests := []struct {
		name    string
		A, B, K int
		want    string
	}{
		{"入力例1", 2, 3, 3, "0 2"},
		{"入力例2", 500000000000, 500000000000, 1000000000000, "0 0"},
		{"", 2, 3, 1, "1 3"},
		{"", 2, 3, 100, "0 0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC149Bその1(tt.A, tt.B, tt.K)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
