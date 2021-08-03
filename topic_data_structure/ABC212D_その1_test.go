package topic_data_structure

import (
	"container/heap"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

type intHeapABC212D []int

func (h intHeapABC212D) Len() int {
	return len(h)
}

func (h intHeapABC212D) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h intHeapABC212D) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeapABC212D) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeapABC212D) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// [ABC212D - Querying Multiset](https://atcoder.jp/contests/abc212/tasks/abc212_d)
func AnswerABC212Dその1(Q int, queries [][]int) string {
	h := make(intHeapABC212D, 0, Q)
	chokin := 0

	out := make([]string, 0, Q)

	for _, query := range queries {
		p := query[0]

		switch p {
		case 1:
			// 新たな追加される値は、それまでの貯金の恩恵がないので差っ引く
			heap.Push(&h, query[1]-chokin)
		case 2:
			// 全体に加算は貯金で表現
			chokin += query[1]
		case 3:
			// 取り出すときは、貯金を適用する
			ball := heap.Pop(&h).(int) + chokin

			out = append(out, strconv.Itoa(ball))
		}
	}

	return strings.Join(out, "\n")
}

func TestAnswerABC212Dその1(t *testing.T) {
	tests := []struct {
		name    string
		Q       int
		queries [][]int
		want    string
	}{
		{
			"sample",
			5,
			[][]int{
				{1, 2},
				{2, 3},
				{1, 7},
				{3},
				{3},
			},
			"5\n7",
		},
		{
			"ex1",
			5,
			[][]int{
				{1, 3},
				{1, 5},
				{3},
				{2, 2},
				{3},
			},
			"3\n7",
		},
		{
			"ex2",
			6,
			[][]int{
				{1, 1000000000},
				{2, 1000000000},
				{2, 1000000000},
				{2, 1000000000},
				{2, 1000000000},
				{3},
			},
			"5000000000",
		},
		{
			"オレオレサンプル",
			7,
			[][]int{
				{1, 8},
				{1, 3},
				{3},
				{3},
				{1, 9},
				{3},
			},
			"3\n8\n9",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC212Dその1(tt.Q, tt.queries)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
