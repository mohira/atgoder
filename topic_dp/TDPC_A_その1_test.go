package topic_dp

import (
	"testing"
)

// [TDPC A - コンテスト](https://atcoder.jp/contests/tdpc/tasks/tdpc_contest)
func AnswerTDPC問題Aその1(N int, P []int) int {
	return 0
}

func TestAnswerTDPC問題Aその1(t *testing.T) {

	tests := []struct {
		name string
		N    int
		P    []int
		want int
	}{
		{"入力例1", 3, []int{2, 3, 5}, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerTDPC問題Aその1(tt.N, tt.P)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}

		})
	}
}
