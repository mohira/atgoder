package ABC_problemB

import (
	"testing"
)

// [ABC079B - Lucas Number](https://atcoder.jp/contests/abc079/tasks/abc079_b)
func AnswerABC079Bその1(N int) int {
	lucasNums := make([]int, 87)
	lucasNums[0] = 2
	lucasNums[1] = 1

	for i := 2; i < len(lucasNums); i++ {
		lucasNums[i] = lucasNums[i-1] + lucasNums[i-2]
	}

	return lucasNums[N]
}

func TestAnswerABC079Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{"入力例1", 5, 11},
		{"入力例2", 86, 939587134549734843},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC079Bその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
