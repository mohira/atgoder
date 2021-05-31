package problemC

import (
	"strconv"
	"strings"
	"testing"
)

// [ABC134C - Exception Handling](https://atcoder.jp/contests/abc134/tasks/abc134_c)
func AnswerABC134Cその1(N int, A []int) string {
	// 答えは、最大値もしくは2番目に大きい値にしかならない
	// 最大値と2番目に大きい値は同じ値の可能性があるのでIdxも調べる
	var maxIdx, maxA int

	for i, a := range A {
		if maxA < a {
			maxIdx = i
			maxA = a
		}
	}

	var secondMaxA int
	for i, a := range A {
		if i == maxIdx {
			continue
		}
		if secondMaxA < a {
			secondMaxA = a
		}
	}

	var s = make([]string, N)

	for i := range A {
		if i == maxIdx {
			s[i] = strconv.Itoa(secondMaxA)
		} else {
			s[i] = strconv.Itoa(maxA)
		}
	}

	return strings.Join(s, "\n")
}

func TestAnswerABC134Cその1(t *testing.T) {
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
			got := AnswerABC134Cその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
