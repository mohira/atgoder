package ARC_problemA

import (
	"atgoder/lib"
	"testing"
)

// [ARC041A - 大好き高橋君](https://atcoder.jp/contests/arc041/tasks/arc041_1)
func AnswerARC041Aその1(x, y, k int) int {
	omote := x

	omote += lib.Min(k, y)
	k = lib.Max(0, k-y)

	omote -= lib.Min(k, x)
	k = lib.Max(0, k-x)

	return omote
}

func TestAnswerARC041Aその1(t *testing.T) {
	tests := []struct {
		name string
		x, y int
		k    int
		want int
	}{
		{"入力例1", 3, 2, 1, 4},
		{"入力例2", 3, 2, 4, 3},
		{"入力例3", 3, 2, 5, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerARC041Aその1(tt.x, tt.y, tt.k)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
