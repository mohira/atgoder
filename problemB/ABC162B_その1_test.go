package problemB

import (
	"testing"
)

// [ABC162B - FizzBuzz Sum](https://atcoder.jp/contests/abc162/tasks/abc162_b)
func AnswerABC162Bその1(N int) int {
	var ans int
	for i := 1; i <= N; i++ {
		isFizzBuzz := (i%3 == 0) || (i%5 == 0)

		if !isFizzBuzz {
			ans += i
		}
	}
	return ans
}

func TestAnswerABC162Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{"入力例1", 15, 60},
		{"入力例2", 1000000, 266666333332},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC162Bその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
