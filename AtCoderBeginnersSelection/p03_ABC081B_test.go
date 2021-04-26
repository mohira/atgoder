package AtCoderBeginnersSelection

import (
	"testing"
)

func isAllEven(numbers []int) bool {
	for _, number := range numbers {
		if number%2 != 0 {
			return false
		}
	}
	return true
}

func toHalf(numbers []int) []int {
	var halfNumbers []int

	for _, a := range numbers {
		halfNumbers = append(halfNumbers, a/2)
	}

	return halfNumbers
}

func AnswerABC081B(N int, A []int) int {
	var opCount int

	for {
		if isAllEven(A) {
			opCount++
			A = toHalf(A)
		} else {
			return opCount
		}
	}

}

func TestAnswerABC081B(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want int
	}{
		{"入力例1", 3, []int{8, 12, 40}, 2},
		{"入力例2", 4, []int{5, 6, 8, 10}, 0},
		{"入力例3", 6, []int{382253568, 723152896, 37802240, 379425024, 404894720, 471526144}, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC081B(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}

		})

	}
}
