package topic_manipulate_digit

import (
	"math"
	"sort"
	"testing"
)

// [ABC192C - Kaprekar Number](https://atcoder.jp/contests/abc192/tasks/abc192_c)
func AnswerABC192Cその1(N, K int) int {
	for i := 0; i < K; i++ {
		N = f(N)
	}
	return N
}

func f(N int) int {
	return g1(N) - g2(N)
}

func g1(n int) int {
	// x を十進法で表したときの各桁の数字を大きい順に並び替えてできる整数
	var ints []int
	for n > 0 {
		ints = append(ints, n%10)
		n /= 10
	}
	sort.Ints(ints)

	result := 0
	for i, x := range ints {
		result += int(math.Pow10(i) * float64(x))
	}

	return result
}

func g2(n int) int {
	// x を十進法で表したときの各桁の数字を小さい順に並び替えてできる整数
	var ints []int
	for n > 0 {
		ints = append(ints, n%10)
		n /= 10
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))

	result := 0
	for i, x := range ints {
		result += int(math.Pow10(i) * float64(x))
	}

	return result
}

func TestAnswerABC192Cその1(t *testing.T) {
	tests := []struct {
		name string
		A, B int
		want int
	}{
		{"入力例1", 314, 2, 693},
		{"入力例2", 1000000000, 100, 0},
		{"入力例3", 6174, 100000, 6174},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC192Cその1(tt.A, tt.B)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
