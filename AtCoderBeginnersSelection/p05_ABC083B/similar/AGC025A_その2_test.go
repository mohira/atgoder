package similar

import (
	"testing"
)

// [AGC025A - Digits Sum](https://atcoder.jp/contests/agc025/tasks/agc025_a)
func AnswerAGC025Aその2(N int) int {
	// 解説を参考したパターン

	// 2<=N なので、べき乗チェックはこれでOK(2の0乗にはならない)
	if N%10 == 0 {
		return 10
	}

	// 10のべき乗でない場合は、各桁の和が下限値
	var ans int

	for N > 0 {
		ans += N % 10
		N /= 10
	}

	return ans
}

func TestAnswerAGC025Aその2(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{"入力例1", 15, 6},
		{"入力例2", 100000, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerAGC025Aその2(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
