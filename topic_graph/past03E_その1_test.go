package topic_graph

import (
	"fmt"
	"strings"
	"testing"
)

// [第三回 アルゴリズム実技検定E - スプリンクラー](https://atcoder.jp/contests/past202005-open/tasks/past202005_e)
func AnswerPAST03Eその1隣接行列を使った解法(N, M, Q int, UV [][]int, C []int, S [][]int) string {
	// 1. 隣接行列をつくる
	// Falseの Nx0 の二次元配列をつくる
	graph := make([][]int, N)

	// 2. M本の辺を受け取る
	for i := 0; i < M; i++ {
		u, v := UV[i][0], UV[i][1]
		u-- // 頂点番号は-1
		v-- // 頂点番号は-1

		// 頂点u から 頂点v への辺がある
		graph[u] = append(graph[u], v)

		// 頂点u から 頂点v への辺がある
		graph[v] = append(graph[v], u)
	}

	var ansString = make([]string, 0, Q)

	// 3. Q個のクエリの処理
	for _, query := range S {
		// スプリンクラーを起動する
		if query[0] == 1 {
			x := query[1]
			x--

			// 頂点の色を登録
			ansString = append(ansString, fmt.Sprintf("%d", C[x]))

			// 隣接する頂点の色を上書きする
			for _, nodeId := range graph[x] {
				C[nodeId] = C[x]
			}

		}

		if query[0] == 2 {
			x := query[1]
			x--
			y := query[2]

			// 頂点の色を登録
			ansString = append(ansString, fmt.Sprintf("%d", C[x]))

			C[x] = y
		}

	}

	return strings.Join(ansString, "\n")
}

func TestPAST03Eその1(t *testing.T) {
	tests := []struct {
		name    string
		N, M, Q int
		UV      [][]int
		C       []int
		S       [][]int
		want    string
	}{
		{
			"入力例1",
			3, 2, 3,
			[][]int{
				{1, 2},
				{2, 3},
			},
			[]int{5, 10, 15},
			[][]int{
				{1, 2},
				{2, 1, 20},
				{1, 1},
			},
			"10\n10\n20",
		},
		{
			"入力例2",
			30, 10, 20,
			[][]int{
				{11, 13},
				{30, 14},
				{6, 4},
				{7, 23},
				{30, 8},
				{17, 4},
				{6, 23},
				{24, 18},
				{26, 25},
				{9, 3},
			},
			[]int{18, 4, 36, 46, 28, 16, 34, 19, 37, 23, 25, 7, 24, 16, 17, 41, 24, 38, 6, 29, 10, 33, 38, 25, 47, 8, 13, 8, 42, 40},
			[][]int{
				{2, 1, 9},
				{1, 8},
				{1, 9},
				{2, 20, 24},
				{2, 26, 18},
				{1, 20},
				{1, 26},
				{2, 24, 31},
				{1, 4},
				{2, 21, 27},
				{1, 25},
				{1, 29},
				{2, 10, 14},
				{2, 2, 19},
				{2, 15, 36},
				{2, 28, 6},
				{2, 13, 5},
				{1, 12},
				{1, 11},
				{2, 14, 43},
			},
			"18\n19\n37\n29\n8\n24\n18\n25\n46\n10\n18\n42\n23\n4\n17\n8\n24\n7\n25\n16",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerPAST03Eその1隣接行列を使った解法(tt.N, tt.M, tt.Q, tt.UV, tt.C, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}

}
