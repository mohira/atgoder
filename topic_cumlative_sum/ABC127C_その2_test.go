package topic_cumlative_sum

import (
	"math"
	"testing"
)

// [ABC127C - Prison](https://atcoder.jp/contests/abc127/tasks/abc127_c)
func AnswerABC127Cその2(N, M int, LR [][]int) int {
	// Notいもす法
	// 範囲が連続なので、もっとも強い制約だけをさがせばいい
	// こんな感じの制約が3つあったとすると
	// |-------------|
	//      |-----------|
	//    |-------|
	//       ***** ← この範囲が答え
	var lMax = 0
	var rMin = math.MaxInt32

	for _, lr := range LR {
		l, r := lr[0], lr[1]

		if lMax < l {
			lMax = l
		}

		if r < rMin {
			rMin = r
		}
	}

	// 1枚も通れない場合があるので注意
	return max(rMin-lMax+1, 0)
}

func max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func TestAnswerABC127Cその2(t *testing.T) {
	tests := []struct {
		name string
		N, M int
		LR   [][]int
		want int
	}{
		{"入力例1",
			4, 2,
			[][]int{
				{1, 3},
				{2, 4},
			},
			2},
		{"入力例2",
			10, 3,
			[][]int{
				{3, 6},
				{5, 7},
				{6, 9},
			},
			1},
		{"入力例3",
			100000, 1,
			[][]int{
				{1, 100000},
			},
			100000},
		{"1枚も通れない場合",
			4, 2,
			[][]int{
				{1, 2},
				{3, 4},
			},
			0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC127Cその2(tt.N, tt.M, tt.LR)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
