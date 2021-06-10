package ABC_problemB

import (
	"atgoder/lib"
	"testing"
)

// [ABC182B - Almost GCD](https://atcoder.jp/contests/abc182/tasks/abc182_b)
func AnswerABC182Bその1(N int, A []int) int {
	var maxGCD int
	var ans int

	for k := 2; k <= lib.FindMax(A); k++ {
		gcd := 0

		for j := 0; j < N; j++ {
			if A[j]%k == 0 {
				gcd++
			}
		}

		if maxGCD < gcd {
			maxGCD = gcd
			ans = k
		}
	}

	return ans
}

func TestAnswerABC182Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want int
	}{
		{"入力例1", 3, []int{3, 12, 7}, 3},
		{"入力例2", 5, []int{8, 9, 18, 90, 72}, 2},             // 3 や 9 でもOK
		{"入力例3", 5, []int{1000, 1000, 1000, 1000, 1000}, 2}, // 1000 でもOK
		{"入力例4", 6, []int{472, 236, 767, 35, 531, 661}, 59},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC182Bその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
