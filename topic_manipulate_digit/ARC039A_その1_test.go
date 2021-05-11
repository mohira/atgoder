package topic_manipulate_digit

import (
	"strconv"
	"strings"
	"testing"
)

// [AR039A - A - B problem](https://aatcoder.jp/contests/arc039/tasks/arc039_a)
func AnswerARC039Aその1(A, B int) int {
	bestA := convertMax(A)
	bestB := convertMin(B)

	diffA := bestA - A
	diffB := B - bestB

	if diffB <= diffA {
		return bestA - B
	} else {
		return A - bestB
	}
}

func toDigits(N int) []int {
	digits := make([]int, 0, 3)

	for _, b := range strings.Split(strconv.Itoa(N), "") {
		n, _ := strconv.Atoi(b)
		digits = append(digits, n)
	}
	return digits
}

func convertMax(A int) int {
	var bestA int
	digitsA := toDigits(A)

	if digitsA[0] < 9 {
		// 100の位の改善
		bestA = 100*9 + 10*digitsA[1] + digitsA[2]
	}

	if digitsA[0] == 9 && digitsA[1] < 9 {
		// 10の位の改善
		bestA = 100*digitsA[0] + 10*9 + digitsA[2]
	}

	if digitsA[0] == 9 && digitsA[1] == 9 && digitsA[2] < 9 {
		// 1の位の改善
		bestA = 100*digitsA[0] + 10*digitsA[1] + 1*9
	}

	return bestA
}

func convertMin(B int) int {
	var bestB int
	digitsB := toDigits(B)

	if 1 < digitsB[0] {
		// 100の位の改善
		bestB = 100*1 + 10*digitsB[1] + digitsB[2]
	}

	if 1 == digitsB[0] && 1 <= digitsB[1] {
		// 10の位の改善
		bestB = 100*digitsB[0] + 10*0 + digitsB[2]
	}

	if 1 == digitsB[0] && 0 == digitsB[1] {
		// 1の位の改善
		bestB = 100*digitsB[0] + 10*digitsB[1] + 0
	}

	return bestB
}

func TestAnswerARC039Aその1(t *testing.T) {
	tests := []struct {
		name string
		A, B int
		want int
	}{
		{"入力例1", 567, 234, 733},
		{"入力例2", 999, 100, 899},
		{"入力例3", 100, 999, -99},
		{"入力例: Bを改善したほうがお得", 934, 999, 735},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerARC039Aその1(tt.A, tt.B)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
