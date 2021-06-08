package topic_dfs

import (
	"container/list"
	"fmt"
)

// [よくやる再帰関数の書き方 〜 n 重 for 文を機械的に 〜 - けんちょんの競プロ精進記録](https://drken1215.hatenablog.com/entry/2020/05/04/190252)
func dfsTemplate(A *list.List, N, M int) {
	if A.Len() == N {
		// やりたい処理: 出力したり、総和をとったり
		pprint(A)
		return
	}

	for v := 0; v < M; v++ {
		// 1. 末尾にvを追加 == 1つ下の階層に探索を進める
		A.PushBack(v)

		// 2. dfsを再帰呼び出し
		dfsTemplate(A, N, M)

		// 3. 末尾の要素を削除(pop) == 親ノードに戻る
		A.Remove(A.Back())
	}
	fmt.Println("========")
}

func pprint(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	fmt.Println()
}

func main() {
	fmt.Println()
	A := list.New()
	N := 5
	M := 3
	dfsTemplate(A, N, M)
}
