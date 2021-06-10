package ABC_problemB

import (
	"strconv"
	"testing"
)

// [ABC126B - YYMM or MMYY](https://atcoder.jp/contests/abc126/tasks/abc126_b)
func AnswerABC126Bその1(S string) string {
	left, right := S[:2], S[2:]

	leftIsYY, rightIsYY := okYY(left), okYY(right)
	leftIsMM, rightIsMM := okMM(left), okMM(right)

	isYYMM := leftIsYY && rightIsMM
	isMMYY := leftIsMM && rightIsYY

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

func okYY(yy string) bool {
	n, _ := strconv.Atoi(yy)

	if 0 <= n && n <= 99 {
		return true
	} else {
		return false
	}
}

func okMM(mm string) bool {
	n, _ := strconv.Atoi(mm)
	if 1 <= n && n <= 12 {
		return true
	} else {
		return false
	}
}

func TestAnswerABC126Bその1(t *testing.T) {
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
			got := AnswerABC126Bその1(tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
