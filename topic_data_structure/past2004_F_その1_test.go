package topic_data_structure

import (
	"container/heap"
	"strconv"
	"strings"
	"testing"
)

type taskHeap []int

func (h taskHeap) Len() int {
	return len(h)
}

func (h taskHeap) Less(i, j int) bool {
	return h[i] > h[j] // 最大値が根
}

func (h taskHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *taskHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *taskHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

// [PAST2004F - タスクの消化](https://atcoder.jp/contests/past202004-open/tasks/past202004_f)
func AnswerPAST2004Fその1(N int, AB [][]int) string {
	// i日目に実行可能なタスクを詰め込む
	taskBucket := make(map[int][]int, N)
	for i := 0; i < N; i++ {
		a, b := AB[i][0], AB[i][1]
		taskBucket[a] = append(taskBucket[a], b)
	}

	h := make(taskHeap, 0, N)
	heap.Init(&h)

	out := make([]string, N)
	currentPoint := 0

	for i := 1; i <= N; i++ { // 1-indexed

		// i日目までに実行可能なタスクを都度追加する
		// 都度の追加なので計算量は、ならしO(N)
		if tasks, ok := taskBucket[i]; ok {
			for _, taskPoint := range tasks {
				heap.Push(&h, taskPoint)
			}
		}

		point := heap.Pop(&h).(int)

		currentPoint += point

		out[i-1] = strconv.Itoa(currentPoint)
	}

	return strings.Join(out, "\n")
}

func TestAnswerPAST2004Fその1(t *testing.T) {
	type args struct {
		N  int
		AB [][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"入力例1",
			args{
				N: 3,
				AB: [][]int{
					{1, 3},
					{2, 2},
					{2, 4},
				},
			},
			"3\n7\n9",
		},
		{
			"入力例2",
			args{
				N: 5,
				AB: [][]int{
					{5, 3},
					{4, 1},
					{3, 4},
					{2, 1},
					{1, 5},
				},
			},
			"5\n6\n10\n11\n14",
		},
		{
			"入力例3",
			args{
				N: 6,
				AB: [][]int{
					{1, 8},
					{1, 6},
					{2, 9},
					{3, 1},
					{3, 2},
					{4, 1},
				},
			},
			"8\n17\n23\n25\n26\n27",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnswerPAST2004Fその1(tt.args.N, tt.args.AB); got != tt.want {
				t.Errorf("AnswerPAST2004Fその1() = %v, want %v", got, tt.want)
			}
		})
	}
}
