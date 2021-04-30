package p08_ABC085C

import (
	"fmt"
	"testing"
)

// [ABC112C - Pyramid](https://atcoder.jp/contests/abc112/tasks/arc086_a)

// https://www.dropbox.com/sh/nx3tnilzqz7df8a/AAA8WeL1zPJwWQd1Vs6JFDmRa/ABC112/C?dl=0&subfolder_nav_tracking=1
func AnswerABC112Cその1(N int, XYH [][3]int) string {
	// すべての i について、 hi = H を満たすかどうかを調べる
	// 1 <= Cx, Cy <= 100 なので、全探索で行ける
	// i元の連立方程式を解く感じ
	//
	// 注意すべきは、 hi = 0 となる地点が与えられた場合
	// まず、条件を満たすピラミッドは**ただ1つ**しか存在しない
	// もし、 hi = 0 なら、未調査の座標を中心とする高さ1のピラミッドが成立してしまうから。
	for Cx := 0; Cx <= 100; Cx++ {
		for Cy := 0; Cy <= 100; Cy++ {
			var CxCyH [][3]int
			var H int
			for _, xyh := range XYH {
				xi, yi, hi := xyh[0], xyh[1], xyh[2]

				// hi >= 1 が重要だ！！
				// https://img.atcoder.jp/abc112/editorial.pdf を読む
				if hi >= 1 {
					H = calcH(Cy, Cx, xi, yi, hi)

					CxCyH = append(CxCyH, [3]int{Cx, Cy, H})
					if Cx == 55 && Cy == 88 {
						fmt.Println(xi, yi, hi, CxCyH)
					}
				}

			}

			if allSame(CxCyH) {
				return fmt.Sprintf("%d %d %d", Cx, Cy, H)
			}
		}
	}

	return ""
}

func calcH(Cy, Cx, x, y, h int) int {
	return h + absInt(x-Cx) + absInt(y-Cy)
}

func absInt(n int) int {
	if n > 0 {
		return n
	} else {
		return -n
	}
}

func allSame(CxCyH [][3]int) bool {
	for i := 0; i < len(CxCyH)-1; i++ {
		if CxCyH[i] != CxCyH[i+1] {
			return false
		}
	}

	return true
}

func TestAnswerABC112Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		XYH  [][3]int
		want string
	}{
		{"入力例1",
			4,
			[][3]int{
				{2, 3, 5},
				{2, 1, 5},
				{1, 2, 5},
				{3, 2, 5},
			},
			"2 2 6"},

		{"入力例2",
			2,
			[][3]int{
				{0, 0, 100},
				{1, 1, 98},
			},
			"0 0 100"},
		{"入力例3",
			3,
			[][3]int{
				{99, 1, 191},
				{100, 1, 192},
				{99, 0, 192},
			},
			"100 0 193"},
		// https://www.dropbox.com/sh/nx3tnilzqz7df8a/AAA8WeL1zPJwWQd1Vs6JFDmRa/ABC112/C?dl=0&subfolder_nav_tracking=1
		{"in04.txt: hi=0 を含んでいるパターン",
			4,
			[][3]int{
				{59, 3, 0},
				{64, 94, 56},
				{5, 62, 11},
				{3, 93, 14},
			},
			"55 80 79"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC112Cその1(tt.N, tt.XYH)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
