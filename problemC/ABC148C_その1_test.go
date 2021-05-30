package topic_amari

import (
	"atgoder/lib"
	"testing"
)

// [ABC148C - Snack](https://atcoder.jp/contests/abc148/tasks/abc148_c)
func AnswerABC148Cその1(A, B int) int {
	return lib.LCM([]int{A, B})
}

func TestAnswerABC148Cその1(t *testing.T) {
	tests := []struct {
		name string
		A, B int
		want int
	}{
		{"入力例1", 2, 3, 6},
		{"入力例2", 123, 456, 18696},
		{"入力例3", 100000, 99999, 9999900000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC148Cその1(tt.A, tt.B)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
