package ABC_problemC_gray

import (
	"testing"
)

// [ABC136C - Build Stairs](https://atcoder.jp/contests/abc136/tasks/abc136_c)
func AnswerABC136Cその1(N int, H []int) string {
	// 後ろから見ていく方が、前から見ていくよりも考えやすい
	for i := N - 1; i > 0; i-- {
		right := H[i]
		left := H[i-1]

		if left <= right {
			continue
		}

		if left-1 <= right {
			H[i-1]--
			continue
		}

		return "No"
	}

	return "Yes"
}

func TestAnswerABC136Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		H    []int
		want string
	}{
		{"入力例1", 5, []int{1, 2, 1, 1, 3}, "Yes"},
		{"入力例2", 4, []int{1, 3, 2, 1}, "No"},
		{"入力例3", 5, []int{1, 2, 3, 4, 5}, "Yes"},
		{"入力例4", 1, []int{1000000000}, "Yes"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC136Cその1(tt.N, tt.H)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
