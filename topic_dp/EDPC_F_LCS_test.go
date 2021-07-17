package topic_dp

import (
	"testing"
)

// [EDPC F - LCS](https://atcoder.jp/contests/dp/tasks/dp_f)
func AnswerEDPC問題Fその1(s, t string) string {
	return ""
}

func TestAnswerEDPC問題Fその1(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{
			"入力例1",
			"axyb",
			"abyxb",
			"axb",
		},
		{
			"入力例2",
			"aa",
			"xayaz",
			"aa",
		},
		{
			"入力例3",
			"a",
			"z",
			"",
		},
		{
			"入力例4",
			"abracadabra",
			"avadakedavra",
			"aaadara",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnswerEDPC問題Fその1(tt.s, tt.t); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
