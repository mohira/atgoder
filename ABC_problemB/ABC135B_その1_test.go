package ABC_problemB

import (
	"testing"
)

// [ABC135B - 0 or 1 Swap](https://atcoder.jp/contests/abc135/tasks/abc135_b)
func AnswerABC135Bその1(N int, P []int) string {
	// Pは {1, 2, ..., N} を並び替えたもの という制約条件見落としてた
	if isAsc(P) {
		return "YES"
	}

	var tmp = make([]int, N)

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			copy(tmp, P)

			tmp[i], tmp[j] = tmp[j], tmp[i]

			if isAsc(tmp) {
				return "YES"
			}
		}
	}

	return "NO"
}

func isAsc(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		now, next := nums[i], nums[i+1]

		if !(now <= next) {
			return false
		}
	}

	return true
}

func TestAnswerABC135Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		P    []int
		want string
	}{
		{"入力例1", 5, []int{5, 2, 3, 4, 1}, "YES"},
		{"入力例2", 5, []int{2, 4, 3, 5, 1}, "NO"},
		{"入力例3", 7, []int{1, 2, 3, 4, 5, 6, 7}, "YES"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC135Bその1(tt.N, tt.P)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
