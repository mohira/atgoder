package ABC_problemB

import (
	"testing"
)

// [ABC123B - Five Dishes](https://atcoder.jp/contests/abc123/tasks/abc123_b)
func AnswerABC123Bその1(A, B, C, D, E int) int {
	// 最後に注文すべきものを探す発想(解説22と同じ)
	dishes := []int{A, B, C, D, E}

	var minDigit = 9
	var minValueIdx int
	for i, dish := range dishes {
		d := dish % 10
		if d != 0 && d < minDigit {
			minDigit = d
			minValueIdx = i
		}
	}

	total := 0
	for i, dish := range dishes {
		if i == minValueIdx {
			continue
		}
		d := dish % 10
		total += dish + (10-d)%10
	}

	return total + dishes[minValueIdx]
}

func TestAnswerABC123Bその1(t *testing.T) {
	tests := []struct {
		name          string
		A, B, C, D, E int
		want          int
	}{
		{"入力例1", 29, 20, 7, 35, 120, 215},
		{"入力例2", 101, 86, 119, 108, 57, 481},
		{"入力例3 最小1桁が複数でるパターン", 5, 10, 5, 0, 0, 25},
		{"入力例4 すべての下1桁が0", 10, 20, 30, 40, 50, 150},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC123Bその1(tt.A, tt.B, tt.C, tt.D, tt.E)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
