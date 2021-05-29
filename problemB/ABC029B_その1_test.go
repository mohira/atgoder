package problemB

import (
	"strings"
	"testing"
)

// [ABC029B - カキ](https://atcoder.jp/contests/abc029/tasks/abc029_b)
func AnswerABC029Bその1(S []string) int {
	var ans int
	for _, s := range S {
		if strings.Contains(s, "r") {
			ans++
		}
	}
	return ans
}

func TestAnswerABC029Bその1(t *testing.T) {
	tests := []struct {
		name string
		S    []string
		want int
	}{
		{
			"入力例1",
			[]string{"january", "february", "march", "april", "may", "june", "july", "august", "september", "october", "november", "december"},
			8,
		},
		{
			"入力例2",
			[]string{"rrrrrrrrrr", "srrrrrrrrr", "rsr", "ssr", "rrs", "srsrrrrrr", "rssrrrrrr", "sss", "rrr", "srr", "rsrrrrrrrr", "ssrrrrrrrr"},
			11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC029Bその1(tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
