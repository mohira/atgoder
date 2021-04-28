package similar

import (
	"testing"
)

// [ABC090B - Some Sums](https://atcoder.jp/contests/abc090/tasks/abc090_b)
func AnswerABC090Bその1(A, B int) int {
	var ans int

	for i := A; i <= B; i++ {
		if isPalindromeNumber(i) {
			ans++
		}
	}

	return ans
}

func isPalindromeNumber(n int) bool {
	// 制約条件より、nは5桁の整数
	// 順番通りの配列 と 逆順にした配列 の比較に持ち込む
	var a, b [5]int

	for i := 0; i < 5; i++ {
		digit := n % 10

		a[i] = digit
		b[4-i] = digit

		n /= 10
	}

	return a == b
}

func TestAnswerABC090Bその1(t *testing.T) {
	tests := []struct {
		name string
		A    int
		B    int
		want int
	}{
		{"入力例1", 11009, 11332, 4},
		{"入力例2", 31415, 92653, 612},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC090Bその1(tt.A, tt.B)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
