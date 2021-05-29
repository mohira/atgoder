package problemB

import (
	"strconv"
	"strings"
	"testing"
)

// [ABC028B - 文字数カウント](https://atcoder.jp/contests/abc028/tasks/abc028_b)
func AnswerABC028Bその1(S string) string {
	bucket := make(map[string]int)

	for _, c := range S {
		bucket[string(c)]++
	}

	var counts = make([]string, 6)
	for i, s := range []string{"A", "B", "C", "D", "E", "F"} {
		counts[i] = strconv.Itoa(bucket[s])
	}

	return strings.Join(counts, " ")
}

func TestAnswerABC028Bその1(t *testing.T) {
	tests := []struct {
		name string
		S    string
		want string
	}{
		{
			"入力例1",
			"BEAF",
			"1 1 0 0 1 1",
		},
		{
			"入力例2",
			"DECADE",
			"1 0 1 2 2 0",
		},
		{
			"入力例3",
			"ABBCCCDDDDEEEEEFFFFFF",
			"1 2 3 4 5 6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC028Bその1(tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
