package problemB

import (
	"testing"
)

// [ABC140B - Buffet](https://atcoder.jp/contests/abc140/tasks/abc140_b)
func AnswerABC140Bその1(N int, A, B, C []int) int {
	var ans int

	// 満足度の合計
	for _, b := range B {
		ans += b
	}

	// ボーナスポイントの計算
	for i := 0; i < N-1; i++ {
		next := A[i+1]
		now := A[i]

		if next-now == 1 {
			ans += C[now-1]
		}
	}

	return ans
}

func TestAnswerABC140Bその1(t *testing.T) {
	tests := []struct {
		name    string
		N       int
		A, B, C []int
		want    int
	}{
		{
			"入力例1",
			3,
			[]int{3, 1, 2},
			[]int{2, 5, 4},
			[]int{3, 6},
			14,
		},
		{
			"入力例2",
			4,
			[]int{2, 3, 4, 1},
			[]int{13, 5, 8, 24},
			[]int{45, 9, 15},
			74,
		},
		{
			"入力例3",
			2,
			[]int{1, 2},
			[]int{50, 50},
			[]int{50},
			150,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC140Bその1(tt.N, tt.A, tt.B, tt.C)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
