package problemB

import (
	"strconv"
	"testing"
)

// [ABC198B - Palindrome with leading zeros](https://atcoder.jp/contests/abc198/tasks/abc198_b)
func AnswerABC198Bその1(N int) string {
	for i := 0; i <= 9; i++ {
		zeros := ""
		for j := 0; j < i; j++ {
			zeros += "0"
		}

		if isPalindrome(zeros + strconv.Itoa(N)) {
			return "Yes"
		}
	}

	return "No"
}

func TestAnswerABC198Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want string
	}{
		{"入力例1", 1210, "Yes"},
		{"入力例2", 777, "Yes"},
		{"入力例3", 123456789, "No"},
		{"最大", 1_000_000_000, "Yes"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC198Bその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
