package ABC_problemC_gray

import (
	"atgoder/lib"
	"container/list"
	"testing"
)

// [ABC188C -  ABC Tournament](https://atcoder.jp/contests/abc188/tasks/abc188_c)
func AnswerABC188Cその3(N int, A []int) int {
	// キューで再現するアイデア
	// 前方から2つとって大きい方を末尾に追加していく
	// ex: 1 4 2 5
	// [1 4]  2 5
	//       [2 5] 4
	//             4 5
	l := list.New()
	for _, a := range A {
		l.PushBack(a)
	}

	for l.Len() > 2 {
		// Pop操作の代わり
		p1, _ := l.Remove(l.Front()).(int)
		p2, _ := l.Remove(l.Front()).(int)

		if p1 < p2 {
			l.PushBack(p2)
		} else {
			l.PushBack(p1)
		}
	}

	front, _ := l.Front().Value.(int)
	back, _ := l.Back().Value.(int)
	secondValue := lib.Min(front, back)

	for i, a := range A {
		if a == secondValue {
			return i + 1
		}
	}

	return 0
}

func TestAnswerABC188Cその3(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want int
	}{
		{"入力例1", 2, []int{1, 4, 2, 5}, 2},
		{"入力例2", 2, []int{3, 1, 5, 4}, 1},
		{"入力例3", 4, []int{6, 13, 12, 5, 3, 7, 10, 11, 16, 9, 8, 15, 2, 1, 14, 4}, 2},
		{"", 1, []int{1000000000, 1}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC188Cその3(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
