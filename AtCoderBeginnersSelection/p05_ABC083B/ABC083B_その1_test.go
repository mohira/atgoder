package similar

import (
	"strconv"
	"testing"
)

// [ABC083B - Some Sums](https://atcoder.jp/contests/abc083/tasks/abc083_b)
func AnswerABC083Bその1(N, A, B int) int {
	// 各桁の和を文字列変換を利用して求めるパターン
	var ans int

	for i := 1; i <= N; i++ {
		sumOfDigits := calcSumOfDigits(i)

		if A <= sumOfDigits && sumOfDigits <= B {
			ans += i
		}
	}

	return ans
}

func calcSumOfDigits(i int) int {
	var sum int

	for _, c := range strconv.Itoa(i) {
		digit, err := strconv.Atoi(string(c))
		if err != nil {
			panic(err)
		}

		sum += digit
	}

	return sum
}

func TestAnswerABC083Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    int
		B    int
		want int
	}{
		{"入力例1", 20, 2, 5, 84},
		{"入力例2", 10, 1, 2, 13},
		{"入力例3", 100, 4, 16, 4554},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC083Bその1(tt.N, tt.A, tt.B)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}

		})

	}
}
