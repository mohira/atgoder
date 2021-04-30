package p06_ABC071B

import (
	"testing"
)

// [ABC047B - Counting Roads](https://atcoder.jp/contests/abc047/tasks/abc047_b)
func AnswerABC047Bその2(W, H, N int, XYA [][3]int) int {
	var xMin, yMin = 0, 0
	var xMax, yMax = W, H

	for _, xya := range XYA {
		x, y, a := xya[0], xya[1], xya[2]

		// x < xi の塗りつぶし
		if a == 1 {
			xMin = max(xMin, x)
		}
		// xi < x の塗りつぶし
		if a == 2 {
			xMax = min(xMax, x)
		}
		// y < yi の塗りつぶし
		if a == 3 {
			yMin = max(yMin, y)
		}
		// yi < y の塗りつぶし
		if a == 4 {
			yMax = min(yMax, y)
		}
	}

	// マイナス x マイナス のときに面積がプラスにならないように注意する
	return max(xMax-xMin, 0) * max(yMax-yMin, 0)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func TestAnswerABC047Bその1(t *testing.T) {
	tests := []struct {
		name string
		W    int
		H    int
		N    int
		XYA  [][3]int
		want int
	}{
		{"入力例1",
			5,
			4,
			2,
			[][3]int{
				{2, 1, 1},
				{3, 3, 4},
			},
			9,
		},
		{"入力例1の条件変え",
			5,
			4,
			2,
			[][3]int{
				{2, 1, 2},
				{3, 3, 3},
			},
			2,
		},
		{"入力例2",
			5,
			4,
			3,
			[][3]int{
				{2, 1, 1},
				{3, 3, 4},
				{1, 4, 2},
			},
			0,
		},
		{"入力例3",
			10,
			10,
			5,
			[][3]int{
				{1, 6, 1},
				{4, 1, 3},
				{6, 9, 4},
				{9, 4, 2},
				{3, 1, 3},
			},
			64,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC047Bその2(tt.W, tt.H, tt.N, tt.XYA)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
