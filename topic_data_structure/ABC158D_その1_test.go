package topic_data_structure

import (
	"container/list"
	"strings"
	"testing"
)

// [ABC158D - String Formation](https://atcoder.jp/contests/abc158/tasks/abc158_d)
func AnswerABC158Dその1(S string, Q int, Query [][]interface{}) string {
	q := list.New()

	// Sは1文字ずつ入れなくてはいけない！
	for _, c := range S {
		q.PushBack(string(c))
	}

	reversed := false

	for i := 0; i < Q; i++ {
		query := Query[i]
		t := query[0]

		if t == 1 {
			reversed = !reversed
			continue
		}

		f, c := query[1], query[2]
		if reversed {
			switch f {
			case 1:
				q.PushBack(c)
			case 2:
				q.PushFront(c)
			}
		} else {
			switch f {
			case 1:
				q.PushFront(c)
			case 2:
				q.PushBack(c)
			}
		}
	}

	// 文字列連結だとTLE
	var out strings.Builder
	out.Grow(q.Len())

	if reversed {
		for e := q.Back(); e != nil; e = e.Prev() {
			out.WriteString(e.Value.(string))
		}
	} else {
		for e := q.Front(); e != nil; e = e.Next() {
			out.WriteString(e.Value.(string))
		}
	}

	return out.String()
}

func TestAnswerABC158Dその1(t *testing.T) {
	tests := []struct {
		name  string
		S     string
		Q     int
		Query [][]interface{}
		want  string
	}{
		{
			"入力例1",
			"a",
			4,
			[][]interface{}{
				{2, 1, "p"},
				{1},
				{2, 2, "c"},
				{1},
			},
			"cpa",
		},
		{
			"入力例2",
			"a",
			6,
			[][]interface{}{
				{2, 2, "a"},
				{2, 1, "b"},
				{1},
				{2, 2, "c"},
				{1},
				{1},
			},
			"aabc",
		},
		{
			"入力例3",
			"y",
			1,
			[][]interface{}{
				{2, 1, "x"},
			},
			"xy",
		},
		{
			"初期文字列の長さが1ではない",
			"XY",
			5,
			[][]interface{}{
				{2, 1, "a"},
				{1},
				{2, 2, "b"},
				{1},
				{1},
			},
			"YXab",
		},
		{
			"WA",
			"byscddbiqvqaotjyqummgdnpqnvt",
			25,
			[][]interface{}{
				{2, 2, "z"},
				{2, 2, "e"},
				{2, 2, "e"},
				{2, 1, "b"},
				{1},
				{2, 2, "w"},
				{2, 2, "f"},
				{2, 1, "l"},
				{2, 1, "e"},
				{2, 2, "u"},
				{2, 1, "g"},
				{1},
				{2, 1, "v"},
				{1},
				{2, 1, "n"},
				{2, 2, "f"},
				{2, 2, "w"},
				{1},
				{1},
				{2, 2, "s"},
				{2, 1, "l"},
				{2, 2, "a"},
				{2, 1, "j"},
				{2, 2, "w"},
				{2, 1, "l"},
			},
			"ljlngeleeztvnqpndgmmuqyjtoaqvqibddcsybbwfuvfwsaw",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC158Dその1(tt.S, tt.Q, tt.Query)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
