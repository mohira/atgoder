package problemB

import (
	"atgoder/lib"
	"testing"
)

// [ABC064B - Traveling AtCoDeer Problem](https://atcoder.jp/contests/abc077/tasks/abc077_b)
func AnswerABC064Bその2(N int, A []int) int {
	return lib.FindMax(A) - lib.FindMin(A)
}

func TestAnswerABC064Bその2(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want int
	}{
		{"入力例1", 4, []int{2, 3, 7, 9}, 7},
		{"入力例2", 8, []int{3, 1, 4, 1, 5, 9, 2, 6}, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC064Bその2(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
