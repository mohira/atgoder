package topic_cumlative_sum

import (
	"math"
	"testing"
)

// [ABC098C - Attention](https://atcoder.jp/contests/abc098/tasks/arc098_a)
func AnswerABC098Cその1(N int, S string) int {
	// 1 2 3 4 5
	// ← → → ← ←
	//     |
	// → → | ← ←
	//     i
	// リーダー(i番目の人)の左右それぞれ調べる
	// | 0 --- i-1 | i | i+1 --- N |
	// | → → → → → |   | ← ← ← ← ← |
	westCumSum := make([]int, 1, N+1) // 「西向きの人」の累積和
	eastCumSum := make([]int, 1, N+1) // 「東向きの人」の累積和

	for i, s := range S {
		if s == 'W' {
			westCumSum = append(westCumSum, westCumSum[i]+1)
			eastCumSum = append(eastCumSum, eastCumSum[i])
		} else {
			westCumSum = append(westCumSum, westCumSum[i])
			eastCumSum = append(eastCumSum, eastCumSum[i]+1)
		}
	}

	var requiredTurns = math.MaxInt32
	for i := 1; i <= N; i++ {
		// i-1番目以前 の 西向きの総数
		a := westCumSum[i-1]

		// i+1番目以降 の 東向きの総数
		b := eastCumSum[N] - eastCumSum[(i+1)-1]

		requiredTurns = min(requiredTurns, a+b)
	}

	return requiredTurns
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func TestAnswerABC098Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    string
		want int
	}{
		{"入力例1", 5, "WEEWW", 1},
		{"入力例2", 12, "WEWEWEEEWWWE", 4},
		{"入力例3", 8, "WWWWWEEE", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC098Cその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
