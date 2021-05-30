package topic_amari

import (
	"testing"
)

// [ABC154C - Distinct or Not](https://atcoder.jp/contests/abc154/tasks/abc154_c)
func AnswerABC154Cその1(N int, A []int) string {
	// バケットで重複を調べるパターン
	bucket := make(map[int]int)

	for _, a := range A {
		bucket[a]++
	}

	for _, freq := range bucket {
		if 2 <= freq {
			return "NO"
		}
	}

	return "YES"
}

func TestAnswerABC154Cその1(t *testing.T) {
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
			got := AnswerABC154Cその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
