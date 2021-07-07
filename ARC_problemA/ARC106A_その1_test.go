package ARC_problemA

import (
	"fmt"
	"testing"
)

// [ARC106A - 106](https://atcoder.jp/contests/arc106/tasks/arc106_a)
func AnswerARC106Aその1(N int) string {
	// 10**18 - 3 ** 38 < 0 => 1<=A<=37
	// 10**18 - 5 ** 26 < 0 => 1<=B<=25
	a := 1
	b := 1

	for i := 1; i < 38; i++ {
		a *= 3
		b = 1
		for j := 1; j < 26; j++ {
			b *= 5
			if N-(a+b) == 0 {
				return fmt.Sprintf("%d %d", i, j)
			}
		}
	}

	return "-1"
}

func TestAnswerARC106Aその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want string
	}{
		{"入力例1", 106, "4 2"},
		{"入力例2", 1024, "-1"},
		{"入力例3", 10460353208, "21 1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerARC106Aその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
