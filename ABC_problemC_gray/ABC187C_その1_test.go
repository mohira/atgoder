package ABC_problemC_gray

import (
	"testing"
)

// [ABC187C - 1-SAT](https://atcoder.jp/contests/abc187/tasks/abc187_c)
func AnswerABC187Cその1(N int, S []string) string {
	// 答えの候補は Si のいずれか
	// Si の時点で、自動的に「!を0文字追加」は満たされる
	// なので、あとは、 「!+Si」が、Sに含まれるかだけを調べればいい
	bucket := make(map[string]int)
	for _, s := range S {
		bucket[s]++
	}

	for _, s := range S {
		if _, ok := bucket["!"+s]; ok {
			return s
		}
	}

	return "satisfiable"
}

func TestAnswerABC187Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		S    []string
		want string
	}{
		{"入力例1", 6, []string{"a", "!a", "b", "!c", "d", "!d"}, "a"},
		{"入力例2", 10, []string{"red", "red", "red", "!orange", "yellow", "!blue", "cyan", "!green", "brown", "!gray"}, "satisfiable"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC187Cその1(tt.N, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
