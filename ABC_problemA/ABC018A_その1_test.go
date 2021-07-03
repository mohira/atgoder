package ABC_problemA

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

// [ABC018A - 豆まき](https://atcoder.jp/contests/abc018/tasks/abc018_a)
func AnswerABC018Aその1(A, B, C int) string {
	bucket := make(map[int]string)

	ints := []int{A, B, C}
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))

	for i, v := range ints {
		bucket[v] = fmt.Sprintf("%d", i+1)
	}

	var out []string
	for _, v := range []int{A, B, C} {
		out = append(out, bucket[v])
	}

	return strings.Join(out, "\n")
}

func TestAnswerABC018Aその1(t *testing.T) {
	tests := []struct {
		name    string
		A, B, C int
		want    string
	}{
		{"入力例1", 12, 18, 11, "2\n1\n3"},
		{"入力例2", 10, 20, 30, "3\n2\n1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC018Aその1(tt.A, tt.B, tt.C)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
