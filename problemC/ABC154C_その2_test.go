package problemC

import (
	"sort"
	"testing"
)

// [ABC154C - Distinct or Not](https://atcoder.jp/contests/abc154/tasks/abc154_c)
func AnswerABC154Cその2(N int, A []int) string {
	// ソートして隣接のダブリをチェックするパターン
	sort.Ints(A)

	for i := 0; i < N-1; i++ {
		if A[i] == A[i+1] {
			return "NO"
		}
	}

	return "YES"
}

func TestAnswerABC154Cその2(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want string
	}{
		{"入力例1", 5, []int{2, 6, 1, 4, 5}, "YES"},
		{"入力例2", 6, []int{4, 1, 3, 1, 6, 2}, "NO"},
		{"入力例3", 2, []int{10000000, 10000000}, "NO"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC154Cその2(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
