package ABC_problemC

import (
	"atgoder/lib"
	"math"
	"testing"
)

// [ABC189C - Mandarin Orange](https://atcoder.jp/contests/abc189/tasks/abc189_c)
func AnswerABC189Cその1(N int, A []int) int {
	// 左端(l)を固定して、右端(r)を伸ばしていく
	// このとき、X(区間中の皿から取るりんごの数)をどう求めるか？
	//  i) ある区間の要素の最小探索 → O(N) でダメ
	// ii) 1つずつ区間を大きくしていく探索 → 常に、2数の比較操作でいける！
	var ans int

	for l := 0; l < N; l++ {
		x := math.MaxInt64 // 初期値

		for r := l; r < N; r++ {
			x = lib.Min(x, A[r]) // その時点でとれるだけ取ることを意味する
			n := r - l + 1       // 現時点での皿の数
			totalApples := x * n // すべての皿からx個数取るので、合計は xn になる

			ans = lib.Max(ans, totalApples)
		}
	}

	return ans
}

func TestAnswerABC189Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want int
	}{
		{"入力例1", 6, []int{2, 4, 4, 9, 4, 9}, 20},
		{"入力例2", 6, []int{200, 4, 4, 9, 4, 9}, 200},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC189Cその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
