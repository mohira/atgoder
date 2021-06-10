package ABC_problemB

import (
	"testing"
)

// [ABC153B - Common Raccoon vs Monster](https://atcoder.jp/contests/abc153/tasks/abc153_b)
func AnswerABC153Bその1(H, N int, A []int) string {
	var amountDamage int
	for _, a := range A {
		amountDamage += a
	}
	if H <= amountDamage {
		return "Yes"
	} else {
		return "No"
	}
}

func TestAnswerABC153Bその1(t *testing.T) {
	tests := []struct {
		name string
		H, N int
		A    []int
		want string
	}{
		{"入力例1", 10, 3, []int{4, 5, 6}, "Yes"},
		{"入力例2", 20, 3, []int{4, 5, 6}, "No"},
		{"入力例3", 210, 5, []int{31, 41, 59, 26, 53}, "Yes"},
		{"入力例4", 211, 5, []int{31, 41, 59, 26, 53}, "No"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC153Bその1(tt.H, tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
