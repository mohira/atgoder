package topic_data_structure

import (
	"container/heap"
	"reflect"
	"testing"
)

type intHeapABC141D []int

func (h intHeapABC141D) Len() int {
	return len(h)
}

func (h intHeapABC141D) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h intHeapABC141D) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeapABC141D) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeapABC141D) Pop() interface{} {
	old := *h

	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

// [ABC141D - Powerful Discount Tickets](https://atcoder.jp/contests/abc141/tasks/abc141_d)
func AnswerABC141Dその1(N, M int, A []int) int {
	h := make(intHeapABC141D, N)
	for i := 0; i < N; i++ {
		h[i] = A[i]
	}

	heap.Init(&h)

	for i := 0; i < M; i++ {
		v := heap.Pop(&h).(int)
		discount := v / 2
		heap.Push(&h, discount)
	}
	var ans int

	for _, v := range h {
		ans += v
	}

	return ans
}

func TestAnswerABC141Dその1(t *testing.T) {
	tests := []struct {
		name string
		N, M int
		A    []int
		want int
	}{
		{
			"入力例1",
			3, 3,
			[]int{2, 13, 8},
			9,
		},
		{
			"入力例2",
			4, 4,
			[]int{1, 9, 3, 5},
			6,
		},
		{
			"入力例3",
			1, 100000,
			[]int{1000000000},
			0,
		},
		{
			"入力例4",
			10, 1,
			[]int{1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000},
			9500000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC141Dその1(tt.N, tt.M, tt.A)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
