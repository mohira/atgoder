package topic_manipulate_digit

import (
	"testing"
)

// [ABC136B - Uneven Numbers](https://atcoder.jp/contests/abc136/tasks/abc136_b)
func AnswerABC136Bその1(N int) int {
	var count int
	for i := 1; i <= N; i++ {
		if getDigitNums(i)%2 == 1 {
			count++
		}
	}
	return count
}

func getDigitNums(N int) int {
	var digit int
	for N > 0 {
		N /= 10
		digit++
	}
	return digit
}

func TestAnswerABC136Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{"入力例1", 11, 9},
		{"入力例2", 136, 46},
		{"入力例3", 100000, 90909},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC136Bその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
