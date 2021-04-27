package similar

import (
	"math/bits"
	"testing"
)

// [ABC068B - Break Number](https://atcoder.jp/contests/abc068/tasks/abc068_b)
func AnswerABC068Bその1(N int) int {
	// 2進数を右から見て、どれだけ0が連続しているかを数える
	var max = -1
	var ans int

	for i := 1; i <= N; i++ {
		count := bits.TrailingZeros(uint(i))

		if max < count {
			max = count
			ans = i
		}
	}
	return ans
}

func TestAnswerABC068Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{"入力例1", 7, 4},
		{"入力例2", 32, 32},
		{"入力例3", 1, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC068Bその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}

		})

	}
}
