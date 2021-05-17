package topic_manipulate_digit

import (
	"testing"
)

// [ABC176B - Multiple of 9](https://atcoder.jp/contests/abc176/tasks/abc176_b)
func AnswerABC176Bその1(N string) string {
	var sum int
	for _, c := range N {
		sum += int(c - '0')
	}

	if sum%9 == 0 {
		return "Yes"
	} else {
		return "No"
	}
}

func TestAnswerABC176Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    string
		want string
	}{
		{"入力例1", "123456789", "Yes"},
		{"入力例2", "0", "Yes"},
		{"入力例3", "31415926535897932384626433832795028841971693993751058209749445923078164062862089986280", "No"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC176Bその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
