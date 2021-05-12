package topic_grid

import (
	"testing"
)

// [ABC197B - Visibility](https://atcoder.jp/contests/abc197/tasks/abc197_b)
func AnswerABC197Bその2(H, W, X, Y int, S []string) int {
	board := toStringBoard(H, W, S)

	// あるマス(i,j)の周囲マスの表現
	// -----------------> y軸
	// | 左上(-1,-1) 真上(-1,0) 右上(-1,1)
	// | 左中( 0,-1) 真中( 0,0) 右中( 0,1)
	// | 左下( 1,-1) 真下( 1,0) 右下( 1,1)
	// ↓
	// x軸

	// 12時から時計回りに差分を書いていく
	var dx = [4]int{-1, 0, 1, 0}
	var dy = [4]int{0, 1, 0, -1}

	X--
	Y--
	var count int

	// (i, j)の周囲4方向に対して、順番に調べていく
	for k := 0; k < 4; k++ {
		x := X
		y := Y

		for {
			// 自身のマスは調べていない
			x += dx[k]
			y += dy[k]

			existsSquare := (0 <= x && x < H) && (0 <= y && y < W)
			if !existsSquare {
				break
			}

			isWall := board[x][y] == "#"
			if isWall {
				break
			}
			count++
		}
	}

	// 自身のマスを数えていないので +1 で補正する
	return count + 1
}

func TestAnswerABC197Bその2(t *testing.T) {
	tests := []struct {
		name       string
		H, W, X, Y int
		S          []string
		want       int
	}{
		{"入力例1",
			4, 4, 2, 2,
			[]string{
				"##..",
				"...#",
				"#.#.",
				".#.#",
			},
			4},

		{"入力例2",
			3, 5, 1, 4,
			[]string{
				"#....",
				"#####",
				"....#",
			},
			4},

		{"入力例3",
			5, 5, 4, 2,
			[]string{
				".#..#",
				"#.###",
				"##...",
				"#..#.",
				"#.###",
			},
			3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC197Bその2(tt.H, tt.W, tt.X, tt.Y, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
