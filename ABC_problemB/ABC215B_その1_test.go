package ABC_problemB

import (
	"testing"
)

// [ABC215B - log2\(N\)](https://atcoder.jp/contests/abc215/tasks/abc215_b)
func AnswerABC215Bその1(N int) int {
	// log計算は精度不足(64bitでは足りないらしい)
	// https://atcoder.jp/contests/abc215/editorial/2480
	// return int(math.Floor(math.Log2(float64(N))))
	var ans int
	x := 1

	for i := 0; i <= 60; i++ {
		if N < x {
			ans = i - 1
			break
		}
		x *= 2
	}

	return ans
}

func TestAnswerABC215Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{"入力例1", 6, 2},
		{"入力例2", 1, 0},
		{"入力例3", 1_000_000_000_000_000_000, 59},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC215Bその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
