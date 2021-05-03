package topic_grid

import (
	"fmt"
	"strconv"
	"testing"
)

// [ABC075B - Minesweeper](https://atcoder.jp/contests/abc075/tasks/abc075_b)
func AnswerABC075Bその1(H, W int, S []string) string {
	// 方針: ゼロパディングして、3x3のたたみこみ

	// 3x3convあともサイズを維持するためにゼロパディングする
	X := prepareZeroPadding(H, W, S)

	Y := conv3x3(H, W, X)

	// 爆弾マスを数値からRuneに書き換える
	for i := 0; i < len(S); i++ {
		for j := 0; j < len(S[i]); j++ {
			if S[i][j] == '#' {
				Y[i][j] = '#'
			}
		}
	}

	// 期待される文字列に変換する
	var ans string

	for i := 0; i < len(Y); i++ {
		for j := 0; j < len(Y[i]); j++ {
			if Y[i][j] == '#' {
				ans += "#"
			} else {
				ans += strconv.Itoa(Y[i][j])
			}
		}
		ans += "\n"
	}

	return ans
}

func prepareZeroPadding(H, W int, S []string) [][]int {
	A := makeEmpty2dSlice(H+2, W+2)

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if S[i][j] == '#' {
				A[i+1][j+1] = 1
			}
		}
	}

	return A
}

func conv3x3(H, W int, paddingArray [][]int) [][]int {
	A := makeEmpty2dSlice(H, W)

	for i := 1; i < H+1; i++ {
		for j := 1; j < W+1; j++ {
			a := paddingArray[i-1][j-1] + paddingArray[i-1][j] + paddingArray[i-1][j+1]
			b := paddingArray[i+0][j-1] + paddingArray[i+0][j] + paddingArray[i+0][j+1]
			c := paddingArray[i+1][j-1] + paddingArray[i+1][j] + paddingArray[i+1][j+1]

			A[i-1][j-1] = a + b + c
		}
	}
	return A
}

func makeEmpty2dSlice(H int, W int) [][]int {
	var A = make([][]int, H)

	for i := 0; i < H; i++ {
		A[i] = make([]int, W)
	}

	return A
}

func pprint(x [][]int) {
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])

	}
	fmt.Println()
}

func TestAnswerABC075Bその1(t *testing.T) {
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
			got := AnswerABC075Bその1(tt.H, tt.W, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
