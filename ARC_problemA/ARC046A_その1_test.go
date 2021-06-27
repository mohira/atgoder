package ARC_problemA

import (
	"atgoder/lib"
	"testing"
)

// [ARC046A - ゾロ目数](https://atcoder.jp/contests/arc046/tasks/arc046_a)
func AnswerARC046Aその1(N int) int {
	var count int

	for i := 1; i <= 555555; i++ {
		if lib.IsZoroNumber(i) {
			count++
		}
		if count == N {
			return i
		}
	}

	return 0
}

func TestAnswerARC046Aその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{"入力例1", 1, 1},
		{"入力例2", 11, 22},
		{"入力例3", 50, 555555},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerARC046Aその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
