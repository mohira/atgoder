package problemC

import (
	"atgoder/lib"
	"testing"
)

// [ABC166C - Peaks](https://atcoder.jp/contests/abc166/tasks/abc166_c)
func AnswerABC166Cその2(N, M int, H []int, AB [][]int) int {
	// 結局、チャンピオンデータに勝てばいいだけ(すべて比較する必要はない)
	// 自分とつながっている展望台の最大の高さが、自身よりも低ければOK
	maxHeightMap := make(map[int]int)
	for i := 0; i < M; i++ {
		a, b := AB[i][0], AB[i][1]
		a--
		b--

		maxHeightMap[a] = lib.Max(maxHeightMap[a], H[b])
		maxHeightMap[b] = lib.Max(maxHeightMap[b], H[a])
	}

	var ans int
	for i := 0; i < N; i++ {
		if maxHeightMap[i] < H[i] {
			ans++
		}
	}

	return ans
}

func TestAnswerABC166Cその2(t *testing.T) {
	tests := []struct {
		name string
		N, M int
		H    []int
		AB   [][]int
		want int
	}{
		{
			"入力例1",
			4, 3,
			[]int{1, 2, 3, 4},
			[][]int{
				{1, 3},
				{2, 3},
				{2, 4},
			},
			2,
		},
		{
			"入力例2",
			6, 5,
			[]int{8, 6, 9, 1, 2, 1},
			[][]int{
				{1, 3},
				{4, 2},
				{4, 3},
				{4, 6},
				{4, 6},
			},
			3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC166Cその2(tt.N, tt.M, tt.H, tt.AB)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
