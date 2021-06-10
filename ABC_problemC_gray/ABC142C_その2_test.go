package ABC_problemC_gray

import (
	"fmt"
	"strings"
	"testing"
)

// [ABC142C - Go to School](https://atcoder.jp/contests/abc142/tasks/abc142_c)
func AnswerABC142Cその2(N int, A []int) string {
	bucket := make([]int, N)
	for i, a := range A {
		a--
		bucket[N-a-1] = i + 1
	}

	// 逆順に出していけばOK
	var s = make([]string, 0, N)
	for i := len(A) - 1; i >= 0; i-- {
		s = append(s, fmt.Sprintf("%d", bucket[i]))
	}

	return strings.Join(s, " ")
}

func TestAnswerABC142Cその2(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want string
	}{
		{"入力例1", 3, []int{2, 3, 1}, "3 1 2"},
		{"入力例1", 5, []int{1, 2, 3, 4, 5}, "1 2 3 4 5"},
		{"入力例1", 8, []int{8, 2, 7, 3, 4, 5, 6, 1}, "8 2 4 5 6 7 3 1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC142Cその2(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
