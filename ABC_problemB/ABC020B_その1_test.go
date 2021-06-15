package ABC_problemB

import (
	"fmt"
	"strconv"
	"testing"
)

// [ABC020B - 足し算](https://atcoder.jp/contests/abc020/tasks/abc020_b)
func AnswerABC020Bその1(A, B int) int {
	s := fmt.Sprintf("%d%d", A, B)
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return n * 2
}

func TestAnswerABC020Bその1(t *testing.T) {
	tests := []struct {
		name string
		A, B int
		want int
	}{
		{"入力例1", 1, 23, 246},
		{"入力例2", 999, 999, 1999998},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC020Bその1(tt.A, tt.B)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
