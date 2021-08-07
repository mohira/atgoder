package topic_data_structure

import (
	"container/heap"
	"reflect"
	"testing"
)

type intHeapABC137D []int

func (h intHeapABC137D) Len() int {
	return len(h)
}

func (h intHeapABC137D) Less(i, j int) bool {
	return h[i] > h[j] // 最大値が根
}

func (h intHeapABC137D) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeapABC137D) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeapABC137D) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

// [ABC137D - Summer Vacation](https://atcoder.jp/contests/abc137/tasks/abc137_d)
func AnswerABC137Dその1(N, M int, AB [][]int) int {
	// 貪欲法: 「真に良い」が確定しないといけない

	// 最終日から初日へ向けて考えていくと、バイトの選択肢が増えていく
	// M  日目       に選べるバイトは、a=0 ← これは存在しない
	// M-1日目(最終日)に選べるバイトは、a=1
	// M-2日目       に選べるバイトは、a=1,2
	// M-3日目       に選べるバイトは、a=1,2,3
	// M-4日目       に選べるバイトは、a=1,2,3,4

	// 右から左に向かって考えていくを図解
	// 入力例2の場合
	// --- day00 day01 day02 day03(M) -> 現在の日付
	//                  +2     |
	//                  +3     |
	//                  +4     |
	//            +1           |
	//            +3           |
	// a <---3-----2-----1-----0-----

	bucket := make(map[int][]int)

	for i := 0; i < N; i++ {
		a, b := AB[i][0], AB[i][1]

		bucket[a] = append(bucket[a], b)
	}

	h := make(intHeapABC137D, 0, N)
	heap.Init(&h)

	var ans int

	for i := 1; i <= M; i++ {
		if rewards, ok := bucket[i]; ok {
			for _, reward := range rewards {
				heap.Push(&h, reward)
			}
		}
		if h.Len() > 0 {
			reward := heap.Pop(&h).(int)
			ans += reward
		}
	}

	return ans
}

func TestAnswerABC137Dその1(t *testing.T) {
	tests := []struct {
		name string
		N, M int
		AB   [][]int
		want int
	}{
		{
			"入力例1",
			3, 4,
			[][]int{
				{4, 3},
				{4, 1},
				{2, 2},
			},
			5,
		},
		{
			"入力例2",
			5, 3,
			[][]int{
				{1, 2},
				{1, 3},
				{1, 4},
				{2, 1},
				{2, 3},
			},
			10,
		},
		{
			"入力例3",
			1, 1,
			[][]int{
				{2, 1},
			},
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC137Dその1(tt.N, tt.M, tt.AB)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
