package problemB

import (
	"testing"
)

// [ABC130B - Bounding](https://atcoder.jp/contests/abc130/tasks/abc130_b)
func AnswerABC130Bその1(N, X int, L []int) int {
	var count int

	for i := 0; i <= N; i++ {
		if calcX(L, i) <= X {
			count++
		}
	}

	return count
}

func calcX(L []int, n int) int {
	if n == 0 {
		return 0
	} else {
		return calcX(L, n-1) + L[n-1]
	}
}

func TestAnswerABC130Bその1(t *testing.T) {
	tests := []struct {
		name string
		N, X int
		L    []int
		want int
	}{
		{"入力例1", 3, 6, []int{3, 4, 5}, 2},
		{"入力例2", 4, 9, []int{3, 3, 3, 3}, 4},
		{"入力例2", 100, 10000, []int{100, 29, 69, 46, 30, 75, 24, 30, 98, 20, 27, 23, 76, 2, 3, 5, 62, 52, 79, 84, 28, 62, 29, 23, 36, 53, 66, 3, 78, 63, 85, 74, 37, 8, 88, 34, 79, 62, 54, 5, 73, 56, 94, 89, 36, 85, 93, 81, 34, 86, 39, 28, 77, 99, 13, 78, 59, 56, 63, 9, 13, 88, 38, 38, 55, 76, 93, 35, 48, 45, 88, 59, 69, 30, 66, 53, 17, 84, 30, 34, 78, 92, 77, 38, 86, 74, 7, 38, 6, 93, 66, 66, 14, 59, 73, 14, 41, 10, 27, 53}, 101},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC130Bその1(tt.N, tt.X, tt.L)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
