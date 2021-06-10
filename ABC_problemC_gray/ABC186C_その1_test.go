package ABC_problemC_gray

import (
	"testing"
)

// [ABC186C - Unlucky 7](https://atcoder.jp/contests/abc186/tasks/abc186_c)
func AnswerABC186Cその1(N int) int {
	var count int
	for i := 1; i <= N; i++ {
		if is8進法で7を含んでいる(i) || is10進法で7を含んでいる(i) {
			continue
		}

		count++
	}

	return count
}

func is8進法で7を含んでいる(n int) bool {
	for n > 0 {
		if n%8 == 7 {
			return true
		}
		n /= 8
	}
	return false
}

func is10進法で7を含んでいる(n int) bool {
	for n > 0 {
		if n%10 == 7 {
			return true
		}
		n /= 10
	}
	return false
}

func TestAnswerABC186Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{"入力例1", 20, 17},
		{"入力例2", 100000, 30555},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC186Cその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
