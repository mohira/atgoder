package problemC

import (
	"atgoder/lib"
	"testing"
)

// [ABC188C -  ABC Tournament](https://atcoder.jp/contests/abc188/tasks/abc188_c)
func AnswerABC188Cその1(N int, A []int) int {
	// 準優勝は、最大値の反対ブロックにおける最大値
	var maxIdx int
	var max int
	for i, a := range A {
		if max < a {
			max = a
			maxIdx = i
		}
	}

	length := len(A)

	var secondA int
	if maxIdx < length/2 {
		secondA = lib.FindMax(A[length/2:])
	} else {
		secondA = lib.FindMax(A[:length/2])
	}

	for i, a := range A {
		if a == secondA {
			return i + 1
		}
	}

	return 0
}

func TestAnswerABC188Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want int
	}{
		{"入力例1", 2, []int{1, 4, 2, 5}, 2},
		{"入力例2", 2, []int{3, 1, 5, 4}, 1},
		{"入力例3", 4, []int{6, 13, 12, 5, 3, 7, 10, 11, 16, 9, 8, 15, 2, 1, 14, 4}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC188Cその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
