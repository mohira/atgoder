package topic_grid

import (
	"testing"
)

// [ABC096C \- Grid Repainting 2](https://atcoder.jp/contests/abc096/tasks/abc096_c)
func AnswerABC096Cその1(H, W int, S []string) string {
	// すべての黒いマスは、上下左右に隣接するマスいずれかが黒でないといけない

	// あるマス(i,j)の周囲4マスの表現
	// -----------------> y軸
	// |            真上(-1,0)
	// | 左中( 0,-1) 真中( 0,0) 右中( 0,1)
	// |            真下( 1,0)
	// ↓
	// x軸
	// 12時から時計回りに差分を書いていく
	var dx = [4]int{-1, 0, 1, 0}
	var dy = [4]int{0, 1, 0, -1}

	var canDraw = true

	canvas := toCanvas(H, W, S)

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			// マス(i,j)が黒マスのとき
			if canvas[i][j] == "#" {
				hasBlackSquareNeighbors := false

				// (i,j)の周囲4マスをチェックする(それぞれ、座標は(x,y))
				for k := 0; k < 4; k++ {
					x := i + dx[k]
					y := j + dy[k]

					// マス(x,y) が存在するかのチェック
					if (0 <= x && x < H) && (0 <= y && y < W) {
						// マス(x,y)が黒マス
						if canvas[x][y] == "#" {
							hasBlackSquareNeighbors = true
						}
					}
				}
				if !hasBlackSquareNeighbors {
					canDraw = false
				}
			}
		}
	}

	if canDraw {
		return "Yes"
	} else {
		return "No"
	}
}

func toCanvas(H int, W int, S []string) [][]string {
	var canvas = make([][]string, H)
	for i := 0; i < H; i++ {
		canvas[i] = make([]string, W)
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			canvas[i][j] = string(S[i][j])
		}
	}

	return canvas
}

func TestAnswerABC096Cその1(t *testing.T) {
	tests := []struct {
		name string
		H    int
		W    int
		S    []string
		want string
	}{
		{"デバッグ: 正方形じゃないと落ちる",
			3,
			2,
			[]string{
				".#",
				".#",
				".#",
			},
			"Yes",
		},
		{"入力例1",
			3,
			3,
			[]string{
				".#.",
				"###",
				".#.",
			},
			"Yes",
		},

		{"入力例2",
			5,
			5,
			[]string{

				"#.#.#",
				".#.#.",
				"#.#.#",
				".#.#.",
				"#.#.#",
			},
			"No",
		},

		{"入力例3",
			11,
			11,
			[]string{
				"...#####...",
				".##.....##.",
				"#..##.##..#",
				"#..##.##..#",
				"#.........#",
				"#...###...#",
				".#########.",
				".#.#.#.#.#.",
				"##.#.#.#.##",
				"..##.#.##..",
				".##..#..##.",
			},
			"Yes",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC096Cその1(tt.H, tt.W, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
