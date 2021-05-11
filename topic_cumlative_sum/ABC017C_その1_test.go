package topic_cumlative_sum

import (
	"math"
	"testing"
)

// [ABC017C - ハイスコア](https://atcoder.jp/contests/abc017/tasks/abc017_3)
func AnswerABC017Cその1(N, M int, LRS [][]int) int {
	// 加算するのではなく、全得点からもっとも影響が薄いところを引き算する発想
	// いもす法で区間ごとの得点を算出して、もっとも小さいところをあきらめる(合計から引く)
	// その区間をあきらめる ⇔ その遺跡を探検しない ⇔ 条件を満たす
	//
	// いもす法で累積和をとったときの各要素の値 → その区間を含む遺跡を回ったときに得られる得点
	imos := make([]int, M+1)

	var pointSum int
	for _, lrs := range LRS {
		l, r, s := lrs[0], lrs[1], lrs[2]
		pointSum += s

		l--
		r--
		imos[l] += s
		imos[r+1] -= s
	}

	cumsum := getCumulativeSum(imos)

	var minPoint = math.MaxInt32
	for _, s := range cumsum[1 : M+1] {
		if s < minPoint {
			minPoint = s
		}
	}

	return pointSum - minPoint
}

func TestAnswerABC017Cその1(t *testing.T) {
	tests := []struct {
		name string
		N, M int
		LRS  [][]int
		want int
	}{
		{"入力例1",
			4, 6,
			[][]int{
				{1, 3, 30},
				{2, 3, 40},
				{3, 6, 25},
				{6, 6, 10},
			},
			80},
		{"入力例2",
			2, 7,
			[][]int{
				{1, 3, 90},
				{5, 7, 90},
			},
			180},
		{"入力例3",
			1, 4,
			[][]int{
				{1, 4, 70},
			},
			0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC017Cその1(tt.N, tt.M, tt.LRS)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
