package topic_dfs

import (
	"bytes"
	"container/list"
	"fmt"
	"log"
	"strconv"
	"testing"
)

// [ABC114C - 755](https://atcoder.jp/contests/abc114/tasks/abc114_c)
var count int

func AnswerABC114Cその1(N int) int {
	count = 0
	l := list.New()
	dfs114C(l, N)
	return count
}

func dfs114C(l *list.List, N int) {
	n := toInteger(l)
	if n > N {
		return
	}

	// 含まれている数字に753がそれぞれ1回以上使われている
	if is753Number(l) {
		count++
	}

	// dfsテンプレ
	for _, v := range []int{3, 5, 7} {
		// 1. 要素を追加 == 1つ深いところにいく
		l.PushBack(v)

		// 2. 再帰呼び出し
		dfs114C(l, N)

		// 3. pop == 1つ上に戻る
		l.Remove(l.Back())
	}

}

func is753Number(l *list.List) bool {
	var used3, used5, used7 bool
	for e := l.Front(); e != nil; e = e.Next() {
		i, ok := e.Value.(int)
		if !ok {
			log.Fatalf("おかしい")
		}

		if i == 7 {
			used7 = true
		}
		if i == 5 {
			used5 = true
		}
		if i == 3 {
			used3 = true
		}

	}

	return used7 && used5 && used3
}

// list.Listを整数に変換する
func toInteger(l *list.List) int {
	if l.Len() == 0 {
		return 0
	}

	var out bytes.Buffer
	for e := l.Front(); e != nil; e = e.Next() {
		i, ok := e.Value.(int)
		if !ok {
			return 0
		}
		out.WriteString(fmt.Sprintf("%d", i))
	}

	n, err := strconv.Atoi(out.String())
	if err != nil {
		return 0
	}

	return n
}

func TestAnswerABC114Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{"入力例1", 575, 4},
		{"入力例2", 3600, 13},
		{"入力例3", 999999999, 26484},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC114Cその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
