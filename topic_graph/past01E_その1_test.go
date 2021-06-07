package topic_graph

import (
	"strings"
	"testing"
)

// [第一回 アルゴリズム実技検定E - SNS のログ](https://atcoder.jp/contests/past201912-open/tasks/past201912_e)
func AnswerPAST01Eその1(N, Q int, S [][]int) string {
	// 隣接行列を使って解く
	graph := make([][]bool, N)
	for i := 0; i < N; i++ {
		graph[i] = make([]bool, N)
	}

	for i := 0; i < Q; i++ {
		query := S[i]
		a := query[1]
		a--

		// 1. aがbをフォローする
		if query[0] == 1 {
			b := query[2]
			b--

			graph[a][b] = true
		}

		// 2. フォロー全返し
		if query[0] == 2 {
			// aをフォローしている人
			// 頂点v から 頂点a への辺があるとき
			for v := 0; v < N; v++ {
				if graph[v][a] {
					// 頂点a->頂点v
					graph[a][v] = true
				}
			}
		}

		// 3. フォローフォロー
		// ポイント: 現時点でフォローしているユーザーだけを列挙する(フォローの先を見ないように！)
		if query[0] == 3 {
			// 途中でフォロー関係を書き換えると、フォローのフォローのフォローみたいなのが発生する
			// なので、1人ユーザーごとに、フォローする対象者リストを用意して、
			// あとで書き換える
			toFollow := make([]int, 0, N)

			// すべての頂点を順番に見る。見ている頂点をvとする
			for v := 0; v < N; v++ {

				// a -> v というフォローがあるとき
				// 頂点aから頂点vに辺がある
				if graph[a][v] {
					// フォローのフォローを調べるため、さらにすべての頂点をみる
					// ここは、頂点v(aがフォローしているユーザ)の視点
					for w := 0; w < N; w++ {
						// 自身へのフォローはないので、「wがaではない」条件をつけないといけない
						if graph[v][w] && w != a {
							// フォローのフォロー
							toFollow = append(toFollow, w)
						}
					}
				}
			}

			// 「フォローのフォロー」を行う
			for _, w := range toFollow {
				graph[a][w] = true
			}
		}
	}

	var ans = make([]string, N)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if graph[i][j] {
				ans[i] += "Y"
			} else {
				ans[i] += "N"
			}
		}
	}

	return strings.Join(ans, "\n")
}

// デバッグするときに便利
//func pprint(graph [][]bool) {
//	fmt.Println()
//	for _, bools := range graph {
//		for _, b := range bools {
//			if b {
//				fmt.Print(1)
//			} else {
//				fmt.Print(0)
//			}
//		}
//		fmt.Println()
//	}
//}

func TestPAST01Eその1(t *testing.T) {
	tests := []struct {
		name string
		N, Q int
		S    [][]int
		want string
	}{
		{
			"入力例1",
			6, 7,
			[][]int{
				{1, 1, 2},
				{1, 2, 3},
				{1, 3, 4},
				{1, 1, 5},
				{1, 5, 6},
				{3, 1},
				{2, 6},
			},
			"NYYNYY\nNNYNNN\nNNNYNN\nNNNNNN\nNNNNNY\nYNNNYN",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerPAST01Eその1(tt.N, tt.Q, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}

}
