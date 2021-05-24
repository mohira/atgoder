package problemB

import (
	"testing"
	"time"
)

// [ABC126B - YYMM or MMYY](https://atcoder.jp/contests/abc126/tasks/abc126_b)
func AnswerABC126Bその2(S string) string {
	left, right := S[:2], S[2:]
	_, errYYMM := time.Parse("0601", left+right)
	_, errMMYY := time.Parse("0601", right+left)

	isYYMM := errYYMM == nil
	isMMYY := errMMYY == nil

	switch {
	case isYYMM && isMMYY:
		return "AMBIGUOUS"
	case isYYMM:
		return "YYMM"
	case isMMYY:
		return "MMYY"
	default:
		return "NA"
	}

}

func TestAnswerABC126Bその2(t *testing.T) {
	tests := []struct {
		name string
		S    string
		want string
	}{
		{"入力例1", "1905", "YYMM"},
		{"入力例2", "0112", "AMBIGUOUS"},
		{"入力例3", "1700", "NA"},
		{"MMYY", "1199", "MMYY"},
		{"MMYY", "0100", "MMYY"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC126Bその2(tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
