package topic_grid

import (
	"strconv"
	"testing"
)

// [第4回PAST C - 隣接カウント](https://atcoder.jp/contests/past202010-open/tasks/past202010_c?lang=ja)
func AnswerPAST4Cその1(N, M int, S []string) string {
	// N, M は分かりにくいいのでRenameしている
	H := N
	W := M
	board := toBoard(H, W, S)

	// あるマス(i,j)の周囲マスの表現
	// 軸のとり方はこれがオススメらしい
	// -----------------> y軸
	// | 左上(-1,-1) 真上(-1,0) 右上(-1,1)
	// | 左中( 0,-1) 真中( 0,0) 右中( 0,1)
	// | 左下( 1,-1) 真下( 1,0) 右下( 1,1)
	// ↓
	// x軸

	// 自身を含む周囲9マス
	var dx = [9]int{-1, -1, 0, 1, 1, 1, 0, -1, 0}
	var dy = [9]int{0, 1, 1, 1, 0, -1, -1, -1, 0}

	// 正解用の配列を用意することも考えたけど、処理がかさむのでやめた
	var ans string

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			// (i,j)の周囲9マスをチェックする
			sharpCount := 0
			for k := 0; k < 9; k++ {
				x := i + dx[k]
				y := j + dy[k]

				// マス(x,y) が存在するかのチェック
				if (0 <= x && x < H) && (0 <= y && y < W) {
					if board[x][y] == "#" {
						sharpCount++
					}
				}
			}
			ans += strconv.Itoa(sharpCount)
		}
		ans += "\n"
	}

	return ans
}

func toBoard(H int, W int, S []string) [][]string {
	var board = make([][]string, H)
	for i := 0; i < H; i++ {
		board[i] = make([]string, W)
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			board[i][j] = string(S[i][j])
		}
	}

	return board
}

func TestAnswerPAST4Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		M    int
		A    []string
		want string
	}{
		{"入力例1",
			3,
			4,
			[]string{
				"#.##",
				"..#.",
				"#...",
			},
			"1333\n2433\n1211\n",
		},
		{"入力例2",
			10,
			12,
			[]string{
				"#.##..#...##",
				"#..#..##...#",
				"##.#....##.#",
				".#..###...#.",
				"#..#..#...##",
				"###...#..###",
				".###.#######",
				".#..#....###",
				".#.##..####.",
				".###....#..#",
			},
			"233322331133\n455432343354\n444344443343\n444344332454\n454335431465\n466434554686\n466434445796\n346554457885\n346542135664\n235431134432\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerPAST4Cその1(tt.N, tt.M, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
