package problemC

import (
	"atgoder/lib"
	"math"
	"testing"
)

// [ABC152C - Low Elements](https://atcoder.jp/contests/abc152/tasks/abc152_c)
func AnswerABC152Cその1(N int, P []int) int {
	// 条件を言い換えると、
	//「 P[i] が登場時点で最小であること」
	// なので、各i において、その時点での最小値を保持するスライスを用意して比較すればいい
	minPs := make([]int, N)
	minP := math.MaxInt64
	for i, p := range P {
		minP = lib.Min(minP, p)
		minPs[i] = minP
	}

	count := 0
	for i := 0; i < N; i++ {
		if P[i] <= minPs[i] {
			count++
		}
	}

	return count
}

func TestAnswerABC152Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		P    []int
		want int
	}{
		{"入力例1", 5, []int{4, 2, 5, 1, 3}, 3},
		{"入力例2", 4, []int{4, 3, 2, 1}, 4},
		{"入力例3", 6, []int{1, 2, 3, 4, 5, 6}, 1},
		{"入力例4", 8, []int{5, 7, 4, 2, 6, 8, 1, 3}, 4},
		{"入力例5", 1, []int{1}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC152Cその1(tt.N, tt.P)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
