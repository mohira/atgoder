package ABC_problemC_gray

import (
	"testing"
)

// [ABC100C - *3 or /2](https://atcoder.jp/contests/abc100/tasks/abc100_c)
func AnswerABC100Cその1(N int, A []int) int {
	// 必ず誰か1人が「2で割れる」を負担する必要がある
	// いくら3倍しても、「2で割れる」回数は増えない
	// → 2で割れる回数で決まる
	var ans int

	for _, a := range A {
		countDivide2 := 0 // 2で割れる回数

		for a > 0 {
			if a%2 == 0 {
				a /= 2
				countDivide2++
			} else {
				break
			}
		}

		ans += countDivide2
	}

	return ans
}

func TestAnswerABC100Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want int
	}{
		{"入力例1", 3, []int{5, 2, 4}, 3},
		{"入力例2", 4, []int{631, 577, 243, 199}, 0},
		{"入力例3", 10, []int{2184, 2126, 1721, 1800, 1024, 2528, 3360, 1945, 1280, 1776}, 39},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC100Cその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
