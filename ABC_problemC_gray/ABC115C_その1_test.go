package ABC_problemC_gray

import (
	"atgoder/lib"
	"math"
	"sort"
	"testing"
)

// [ABC115C - Christmas Eve](https://atcoder.jp/contests/abc115/tasks/abc115_c)
func AnswerABC115Cその1(N, K int, H []int) int {
	// 部分集合のうちの「最小の高さの木」を固定した場合、残りは低い順に選んでいくのがベスト
	// ってことは、昇順ソートして全方からK個ずつ選んでいく
	// で、この時点ではO(N)で、残りは部分集合の最大値を求めるわけです
	// 単純に最大値を探すとするとO(N)になってしまうんだが、
	// ソートされていれば、minとmaxはO(1)で求められるのがミソ！
	sort.Ints(H)

	minDiff := math.MaxInt64
	for i := 0; i+K <= N; i++ {
		subHeights := H[i : i+K]
		min := subHeights[0]
		max := subHeights[K-1]

		minDiff = lib.Min(minDiff, max-min)
	}

	return minDiff
}

func TestAnswerABC115Cその1(t *testing.T) {
	tests := []struct {
		name string
		N, K int
		H    []int
		want int
	}{
		{
			"入力例1",
			5, 3,
			[]int{10, 15, 11, 14, 12},
			2,
		},
		{
			"入力例2",
			5, 3,
			[]int{5, 7, 5, 7, 7},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC115Cその1(tt.N, tt.K, tt.H)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
