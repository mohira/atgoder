package topic_next_permutation

import (
	"testing"
)

// [ABC150C - Count Order](https://atcoder.jp/contests/abc150/tasks/abc150_c)
func AnswerABC150Cその1(N int, P, Q []int) int {
	ints := make([]int, 0, N)
	for i := 1; i <= N; i++ {
		ints = append(ints, i)
	}

	var pIndex, qIndex int
	for i, pattern := range permutation(ints) {
		if isSameSlice(pattern, P) {
			pIndex = i
		}

		if isSameSlice(pattern, Q) {
			qIndex = i
		}
	}

	return absInt(pIndex - qIndex)
}

func isSameSlice(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}

func absInt(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

// [Go言語でスライスの要素の順列、組み合わせを与える \- Camera Obscura](http://obelisk.hatenablog.com/entry/2018/10/09/025240)
func remove(ar []int, i int) []int {
	tmp := make([]int, len(ar))
	copy(tmp, ar)
	return append(tmp[0:i], tmp[i+1:]...)
}

// [Go言語でスライスの要素の順列、組み合わせを与える \- Camera Obscura](http://obelisk.hatenablog.com/entry/2018/10/09/025240)
func permutation(intSlice []int) [][]int {
	var result [][]int
	if len(intSlice) == 1 {
		return append(result, intSlice)
	}
	for i, a := range intSlice {
		for _, b := range permutation(remove(intSlice, i)) {
			result = append(result, append([]int{a}, b...))
		}
	}
	return result
}

func TestAnswerABC150Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		P    []int
		Q    []int
		want int
	}{
		{"入力例1",
			3,
			[]int{1, 3, 2},
			[]int{3, 1, 2},
			3},
		{"入力例2",
			8,
			[]int{7, 3, 5, 4, 2, 1, 6, 8},
			[]int{3, 8, 2, 5, 4, 6, 7, 1},
			17517},
		{"入力例3",
			3,
			[]int{1, 2, 3},
			[]int{1, 2, 3},
			0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC150Cその1(tt.N, tt.P, tt.Q)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
