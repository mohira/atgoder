package ARC_problemA

import (
	"testing"
)

// [ARC105A - Fourtune Cookies](https://atcoder.jp/contests/arc105/tasks/arc105_a)
func AnswerARC105Aその1(A, B, C, D int) string {
	n := 4
	cookies := []int{A, B, C, D}

	for bit := 0; bit < (1 << n); bit++ {
		eaten := 0 // 食べたクッキー
		left := 0  // 残ったクッキー

		for i := 0; i < n; i++ {
			if (bit>>i)&1 == 1 {
				eaten += cookies[i]
			} else {
				left += cookies[i]
			}
		}

		if eaten == left {
			return "Yes"
		}

	}
	return "No"
}

func TestAnswerARC105Aその1(t *testing.T) {
	tests := []struct {
		name       string
		A, B, C, D int
		want       string
	}{
		{"入力例1", 1, 3, 2, 4, "Yes"},
		{"入力例2", 1, 2, 4, 8, "No"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerARC105Aその1(tt.A, tt.B, tt.C, tt.D)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
