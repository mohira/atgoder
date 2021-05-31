package problemC

import (
	"atgoder/lib"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// [ABC134C - Exception Handling](https://atcoder.jp/contests/abc134/tasks/abc134_c)
func AnswerABC134Cその2(N int, A []int) string {
	// 答えは、最大値もしくは2番目に大きい値にしかならない
	// N番目に大きい値はソートを使うのが楽
	maxA := lib.FindMax(A)
	secondMaxA := findSecondMax(A)

	var s = make([]string, N)

	for i, a := range A {
		if a == maxA {
			s[i] = strconv.Itoa(secondMaxA)
		} else {
			s[i] = strconv.Itoa(maxA)
		}
	}

	return strings.Join(s, "\n")
}

func findSecondMax(nums []int) int {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	sort.Ints(tmp)

	return tmp[len(tmp)-2]
}

func TestAnswerABC134Cその2(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want string
	}{
		{"入力例1", 3, []int{1, 4, 3}, "4\n3\n4"},
		{"入力例2", 2, []int{5, 5}, "5\n5"},
		{"入力例3", 3, []int{5, 5, 5}, "5\n5\n5"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC134Cその2(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
