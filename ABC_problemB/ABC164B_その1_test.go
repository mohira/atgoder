package ABC_problemB

import (
	"testing"
)

// [ABC164B - Battle](https://atcoder.jp/contests/abc164/tasks/abc164_b)
func AnswerABC164Bその1(A, B, C, D int) string {
	for {
		// 高橋くんの攻撃
		C -= B
		if C <= 0 {
			return "Yes"
		}

		// 青木くんの攻撃
		A -= D
		if A <= 0 {
			return "No"
		}
	}
}

func TestAnswerABC164Bその1(t *testing.T) {
	tests := []struct {
		name       string
		A, B, C, D int
		want       string
	}{
		{"入力例1", 10, 9, 10, 10, "No"},
		{"入力例2", 46, 4, 40, 5, "Yes"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC164Bその1(tt.A, tt.B, tt.C, tt.D)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
