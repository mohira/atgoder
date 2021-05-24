package problemB

import (
	"testing"
)

// [ABC117B - Can you solve this?](https://atcoder.jp/contests/abc117/tasks/abc117_b)
func AnswerABC117Bその1(N int, L []int) string {
	var total = sumInts(L)
	var maxL = findMax(L)

	if maxL < total-maxL {
		return "Yes"
	} else {
		return "No"
	}
}

func TestAnswerABC117Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		L    []int
		want string
	}{
		{"入力例1", 4, []int{3, 8, 5, 1}, "Yes"},
		{"入力例2", 4, []int{3, 8, 4, 1}, "No"},
		{"入力例3", 10, []int{1, 8, 10, 5, 8, 12, 34, 100, 11, 3}, "No"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC117Bその1(tt.N, tt.L)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
