package problemC

import (
	"math"
	"sort"
	"testing"
)

// [ABC138C - Alchemist](https://atcoder.jp/contests/abc138/tasks/abc138_c)
func AnswerABC138Cその1(N int, V []int) float64 {
	// 合成が後になるほど、2で割る影響が少ないので、価値が低い順に合成すればいい
	sort.Ints(V)

	var ans = float64(V[0])
	for _, v := range V[1:] {
		ans = (ans + float64(v)) / 2
	}
	return ans
}

func TestAnswerABC138Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		V    []int
		want float64
	}{
		{"入力例1", 2, []int{3, 4}, 3.5},
		{"入力例2", 3, []int{500, 300, 200}, 375},
		{"入力例3", 5, []int{138, 138, 138, 138, 138}, 138},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC138Cその1(tt.N, tt.V)

			if math.Abs(got-tt.want) > 1e-5 {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
