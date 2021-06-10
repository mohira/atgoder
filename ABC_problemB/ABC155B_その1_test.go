package ABC_problemB

import (
	"testing"
)

// [ABC155B - Papers, Please](https://atcoder.jp/contests/abc155/tasks/abc155_b)
func AnswerABC155Bその1(N int, A []int) string {
	var approved = true

	for _, a := range A {
		if a%2 == 0 {
			if !(a%3 == 0 || a%5 == 0) {
				approved = false
			}
		}
	}

	if approved {
		return "APPROVED"
	} else {
		return "DENIED"
	}
}

func TestAnswerABC155Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want string
	}{
		{"入力例1", 5, []int{6, 7, 9, 10, 31}, "APPROVED"},
		{"入力例2", 3, []int{28, 27, 24}, "DENIED"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC155Bその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
