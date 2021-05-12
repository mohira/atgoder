package topic_grid

import (
	"testing"
)

// [ABC197B - Visibility](https://atcoder.jp/contests/abc197/tasks/abc197_b)
func AnswerABC197Bその1(H, W, X, Y int, S []string) int {
	board := toStringBoard(H, W, S)

	X--
	Y--
	var count int

	// 同じ行を左方向に調べていく
	for j := Y; j >= 0; j-- {
		if board[X][j] == "#" {
			break
		}
		count++
	}

	// 同じ行を右方向に調べていく
	for j := Y; j < W; j++ {
		if board[X][j] == "#" {
			break
		}
		count++
	}

	// 同じ列を上方向に調べていく
	for i := X; i >= 0; i-- {
		if board[i][Y] == "#" {
			break
		}
		count++
	}

	// 同じ列を下方向に調べていく
	for i := X; i < H; i++ {
		if board[i][Y] == "#" {
			break
		}
		count++
	}

	// 自身のマスを4回数えているので -3 で補正する
	return count - 3
}

func toStringBoard(H, W int, S []string) [][]string {
	var board = make([][]string, H)

	for i := 0; i < H; i++ {
		board[i] = make([]string, W)

		for j := 0; j < W; j++ {
			board[i][j] = string(S[i][j])
		}
	}

	return board
}

func TestAnswerABC197Bその1(t *testing.T) {
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
			got := AnswerABC197Bその1(tt.H, tt.W, tt.X, tt.Y, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
