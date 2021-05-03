package p08_ABC049C

import (
	"testing"
)

// [ABC072C - Together](https://atcoder.jp/contests/abc072/tasks/arc072_a)
func AnswerABC072Cその1(N int, A []int) int {
	// 操作後の数: 出現頻度
	var bucket = make(map[int]int)

	for _, a := range A {
		bucket[a-1]++
		bucket[a]++
		bucket[a+1]++
	}

	var maxFrequency int

	for _, f := range bucket {
		if maxFrequency < f {
			maxFrequency = f
		}
	}

	return maxFrequency
}

func TestAnswerABC072Cその2(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want int
	}{
		{"入力例1", 7, []int{3, 1, 4, 1, 5, 9, 2}, 4},
		{"入力例2", 10, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 3},
		{"入力例3", 1, []int{99999}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC072Cその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
