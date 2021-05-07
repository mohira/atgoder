package topic_counter

import (
	"strconv"
	"testing"
)

// [ABC163C - management](https://atcoder.jp/contests/abc163/tasks/abc163_c)
func AnswerABC163Cその1(N int, A []int) string {
	bucket := make(map[int]int)

	for _, a := range A {
		bucket[a]++
	}

	var ans string
	for i := 1; i <= N; i++ {
		ans += strconv.Itoa(bucket[i]) + "\n"
	}
	return ans
}

func TestAnswerABC163Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want string
	}{
		{"入力例1", 5, []int{1, 1, 2, 2}, "2\n2\n0\n0\n0\n"},
		{"入力例2", 10, []int{1, 1, 1, 1, 1, 1, 1, 1, 1}, "9\n0\n0\n0\n0\n0\n0\n0\n0\n0\n"},
		{"入力例3", 7, []int{1, 2, 3, 4, 5, 6}, "1\n1\n1\n1\n1\n1\n0\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC163Cその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
