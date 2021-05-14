package problemB

import (
	"strconv"
	"strings"
	"testing"
)

// [ABC191B - Remove It](https://atcoder.jp/contests/abc191/tasks/abc191_b)
func AnswerABC191Bその1(N, X int, A []int) string {
	var SelectedNums = make([]int, 0, N)

	for _, a := range A {
		if a != X {
			SelectedNums = append(SelectedNums, a)
		}
	}

	var strs = make([]string, len(SelectedNums))
	for i, a := range SelectedNums {
		strs[i] = strconv.Itoa(a)
	}

	return strings.Join(strs, " ")
}

func TestAnswerABC191Bその1(t *testing.T) {
	tests := []struct {
		name string
		N, X int
		A    []int
		want string
	}{
		{"入力例1", 5, 5, []int{3, 5, 6, 5, 4}, "3 6 4"},
		{"入力例2", 3, 3, []int{3, 3, 3}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC191Bその1(tt.N, tt.X, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
