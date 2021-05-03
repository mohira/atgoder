package topic_grid

import (
	"strconv"
	"testing"
)

// [ABC075B - Minesweeper](https://atcoder.jp/contests/abc075/tasks/abc075_b)
func AnswerABC075Bその2(H, W int, S []string) string {
	// ボード変換
	board := convertToBoard(H, W, S)

	// あるマス(i,j)の周囲8マスの表現
	// 軸のとり方はこれがオススメらしい
	// -----------------> y軸
	// | 左上(-1,-1) 真上(-1,0) 右上(-1,1)
	// | 左中( 0,-1) 真中( 0,0) 右中( 0,1)
	// | 左下( 1,-1) 真下( 1,0) 右下( 1,1)
	// ↓
	// x軸

	var dx = [8]int{-1, -1, 0, 1, 1, 1, 0, -1}
	var dy = [8]int{0, 1, 1, 1, 0, -1, -1, -1}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			var bombCount int

			// (i,j)の周囲8マスをチェックする
			// 周囲マスの座標を(x,y)としている
			for k := 0; k < 8; k++ {
				x := i + dx[k]
				y := j + dy[k]

				// マス(x,y) が存在するかのチェック
				if (0 <= x && x < H) && (0 <= y && y < W) {
					if board[x][y] == '#' {
						bombCount++
					}
				}
			}
			if board[i][j] != '#' {
				board[i][j] = bombCount
			}
		}
	}

	return toAnswer(board)
}

func toAnswer(board [][]int) string {
	var ans string

	H := len(board)
	W := len(board[0])

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			v := board[i][j]
			if v == '#' {
				ans += "#"
			} else {
				ans += strconv.Itoa(v)
			}
		}
		ans += "\n"
	}
	return ans
}

func convertToBoard(H, W int, S []string) [][]int {
	var board = make([][]int, H)
	for i := 0; i < H; i++ {
		board[i] = make([]int, W)
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if S[i][j] == '.' {
				board[i][j] = 0
			} else {
				board[i][j] = '#'
			}
		}
	}
	return board
}

func TestAnswerABC075Bその2(t *testing.T) {
	tests := []struct {
		name string
		H    int
		W    int
		S    []string
		want string
	}{
		{"入力例1",
			3,
			5,
			[]string{
				".....",
				".#.#.",
				".....",
			},
			"11211\n1#2#1\n11211\n",
		},
		{"入力例2",
			3,
			5,
			[]string{
				"#####",
				"#####",
				"#####",
			},
			"#####\n#####\n#####\n",
		},

		{"入力例3",
			6,
			6,
			[]string{
				"#####.",
				"#.#.##",
				"####.#",
				".#..#.",
				"#.##..",
				"#.#...",
			},
			"#####3\n#8#7##\n####5#\n4#65#2\n#5##21\n#4#310\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC075Bその2(tt.H, tt.W, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
