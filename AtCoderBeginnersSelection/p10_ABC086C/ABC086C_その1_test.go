package p08_ABC086C

import (
	"testing"
)

// [ABC086C - Traveling](https://atcoder.jp/contests/abc086/tasks/arc089_a)
func AnswerABC086Cその1(N int, TXY [][3]int) string {
	t := []int{0}
	x := []int{0}
	y := []int{0}

	for _, txy := range TXY {
		ti, xi, yi := txy[0], txy[1], txy[2]
		t = append(t, ti)
		x = append(x, xi)
		y = append(y, yi)
	}

	for i := 0; i < N; i++ {
		dt := t[i+1] - t[i]
		dist := abs(x[i+1]-x[i]) + abs(y[i+1]-y[i])

		// 移動時間より遠くへは行けない
		if dt < dist {
			return "No"
		}

		// 移動時間と距離それぞれの偶奇が一致しないなら、いくら時間をかけてもぴったり移動は不可能
		if dt%2 != dist%2 {
			return "No"
		}
	}

	return "Yes"
}

func abs(i int) int {
	if 0 < i {
		return i
	} else {
		return -i
	}
}

func TestAnswerABC086Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		TXY  [][3]int
		want string
	}{
		{"入力例1", 2, [][3]int{
			{3, 1, 2},
			{6, 1, 1},
		}, "Yes"},

		{"入力例2", 1, [][3]int{
			{2, 100, 100},
		}, "No"},

		{"入力例3", 2, [][3]int{
			{5, 1, 1},
			{100, 1, 1},
		}, "No"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC086Cその1(tt.N, tt.TXY)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
