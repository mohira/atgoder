package topic_data_structure

import (
	"container/heap"
	"testing"
)

type Machine struct {
	a     int
	b     int
	count int // 生産量
}

func (m *Machine) produceTime() int {
	return m.a + m.count*m.b
}

type machineHeap []Machine

func (h machineHeap) Len() int {
	return len(h)
}

func (h machineHeap) Less(i, j int) bool {
	return h[i].produceTime() < h[j].produceTime() // 最小値が根
}

func (h machineHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *machineHeap) Push(x interface{}) {
	*h = append(*h, x.(Machine))
}

func (h *machineHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

// [CODE THANKS FESTIVAL 2017C - Factory](https://atcoder.jp/contests/code-thanks-festival-2017-open/tasks/code_thanks_festival_2017_c)
func AnswerCodeThanksFestival2017Cその1(N, K int, AB [][]int) int {
	h := make(machineHeap, N)
	for i := 0; i < N; i++ {
		h[i] = Machine{
			a:     AB[i][0],
			b:     AB[i][1],
			count: 0,
		}
	}
	heap.Init(&h)

	var ans int
	for i := 0; i < K; i++ {
		m := heap.Pop(&h).(Machine)
		ans += m.produceTime()

		heap.Push(&h, Machine{
			a:     m.a,
			b:     m.b,
			count: m.count + 1,
		})
	}

	return ans
}

func TestAnswerCodeThanksFestival2017Cその1(t *testing.T) {
	type args struct {
		N  int
		K  int
		AB [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"入力例1",
			args{
				N: 3,
				K: 3,
				AB: [][]int{
					{1, 3},
					{2, 0},
					{3, 4},
				},
			},
			5,
		},
		{
			"入力例2",
			args{
				N: 10,
				K: 100000,
				AB: [][]int{
					{22, 59},
					{26, 60},
					{72, 72},
					{47, 3},
					{97, 16},
					{75, 41},
					{82, 77},
					{17, 97},
					{32, 32},
					{28, 7},
				},
			},
			7521307799,
		},
		{
			"入力例3",
			args{
				N: 1,
				K: 100000,
				AB: [][]int{
					{1000000000, 1000000000},
				},
			},
			5000050000000000000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnswerCodeThanksFestival2017Cその1(tt.args.N, tt.args.K, tt.args.AB); got != tt.want {
				t.Errorf("AnswerCodeThanksFestival2017Cその1() = %v, want %v", got, tt.want)
			}
		})
	}
}
