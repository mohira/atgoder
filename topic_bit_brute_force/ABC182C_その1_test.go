package topic_bit_brute_force

import (
	"math"
	"strconv"
	"testing"
)

// [ABC182C - To 3](https://atcoder.jp/contests/abc182/tasks/abc182_c)
func AnswerABC182Cその1(N int) int {
	A := toInts(N)
	K := len(A)

	var can bool
	var minDeleteCount = math.MaxInt32

	// 0 はすべてを消すことになるので、 1からスタート
	for bit := 1; bit < (1 << K); bit++ {
		stringN := ""
		deleteCount := 0

		for i := 0; i < K; i++ {
			// フラグが立っているのは残す桁
			if (bit>>i)&1 == 1 {
				stringN += strconv.Itoa(A[i])
			} else {
				deleteCount++
			}
		}

		if newN, _ := strconv.Atoi(stringN); newN%3 == 0 {
			can = true
			if deleteCount < minDeleteCount {
				minDeleteCount = deleteCount
			}
		}
	}

	if can {
		return minDeleteCount
	} else {
		return -1
	}
}

func toInts(N int) []int {
	var A []int

	for N > 0 {
		A = append(A, N%10)
		N /= 10
	}

	return A
}

func TestAnswerABC182Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{"入力例1", 35, 1},
		{"入力例2", 369, 0},
		{"入力例3", 6227384, 1},
		{"入力例4", 11, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC182Cその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
