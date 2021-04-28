package p06_ABC088B

import (
	"sort"
	"testing"
)

// [ABC088B - Card Game for Two](https://atcoder.jp/contests/abc088/tasks/abc088_b)
func AnswerABC088Bその1(N int, A []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(A)))

	var alicePoint int
	for i := 0; i < N; i += 2 {
		alicePoint += A[i]
	}

	var bobPoint int
	for i := 1; i < N; i += 2 {
		bobPoint += A[i]
	}

	return alicePoint - bobPoint
}

func TestAnswerABC088Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want int
	}{
		{"入力例1", 2, []int{3, 1}, 2},
		{"入力例2", 3, []int{2, 7, 4}, 5},
		{"入力例3", 4, []int{20, 18, 2, 18}, 18},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC088Bその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}

		})

	}
}
