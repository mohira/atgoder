package similar

import (
	"testing"
)

// [ABC090B - Some Sums](https://atcoder.jp/contests/abc090/tasks/abc090_b)
func AnswerABC090Bその2(A, B int) int {
	var ans int

	for i := A; i <= B; i++ {
		// 制約条件より、nは5桁の整数なので、
		// 10000の位 == 1の位 かつ  1000の位 == 10の位 を確認すればOK
		// 100の位(3桁目)は回文判定には不要
		digit10000 := (i / 10000) % 10
		digit1000 := (i / 1000) % 10
		digit10 := (i / 10) % 10
		digit1 := (i / 1) % 10

		if (digit10000 == digit1) && (digit1000 == digit10) {
			ans++
		}
	}

	return ans
}

func TestAnswerABC090Bその2(t *testing.T) {
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
			got := AnswerABC090Bその2(tt.A, tt.B)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
