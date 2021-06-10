package ABC_problemC_gray

import (
	"testing"
)

// [ABC164C - gacha](https://atcoder.jp/contests/abc164/tasks/abc164_c)
func AnswerABC164Cその1(N int, S []string) int {
	bucket := make(map[string]int)

	for _, s := range S {
		bucket[s]++
	}

	return len(bucket)
}

func TestAnswerABC164Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		S    []string
		want int
	}{
		{"入力例1", 3, []string{"apple", "orange", "apple"}, 2},
		{"入力例2", 5, []string{"grape", "grape", "grape", "grape", "grape"}, 1},
		{"入力例3", 4, []string{"aaaa", "a", "aaa", "aa"}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC164Cその1(tt.N, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
