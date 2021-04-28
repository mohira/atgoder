package similar

import (
	"testing"
)

// [ABC087B - Coins](https://atcoder.jp/contests/abc087/tasks/abc087_b)
func AnswerABC087Bその1(A, B, C, X int) int {
	var cnt int

	for i := 0; i <= A; i++ {
		for j := 0; j <= B; j++ {
			// X = 500a + 100b + 50c の式変形
			// X は 50 で割り切れる前提
			num50yen := X/50 - 10*i - 2*j

			// cの制約条件の確認
			if 0 <= num50yen && num50yen <= C {
				cnt++
			}
		}
	}

	return cnt
}

func TestAnswerABC087Bその1(t *testing.T) {
	tests := []struct {
		name string
		A    int
		B    int
		C    int
		X    int
		want int
	}{
		{"入力例1", 2, 2, 2, 100, 2},
		{"入力例2", 5, 1, 0, 150, 0},
		{"入力例3", 30, 40, 50, 6000, 213},
		{"O(N^3)だと無理なやつ", 20000, 20000, 20000, 50, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC087Bその1(tt.A, tt.B, tt.C, tt.X)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}

		})

	}
}
