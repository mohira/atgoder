package ABC_problemC

import (
	"sort"
	"strconv"
	"strings"
	"testing"
)

// [ABC036C - 座圧](https://atcoder.jp/contests/abc036/tasks/abc036_c)
func AnswerABC036Cその1(N int, A []int) string {
	bucket := make(map[int]int)
	for _, a := range A {
		bucket[a] = 1
	}

	nums := make([]int, 0, N)
	for num, _ := range bucket {
		nums = append(nums, num)
	}
	sort.Ints(nums)

	for i, num := range nums {
		bucket[num] = i
	}

	out := make([]string, N)
	for i, a := range A {
		v := strconv.Itoa(bucket[a])
		out[i] = v
	}

	return strings.Join(out, "\n")
}

func TestAnswerABC036Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want string
	}{
		{"入力例1", 5, []int{3, 3, 1, 6, 1}, "1\n1\n0\n2\n0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC036Cその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
