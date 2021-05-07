package topic_counter

import (
	"testing"
)

// [ABC082C - Good Sequence](https://atcoder.jp/contests/abc082/tasks/arc087_a)
func AnswerABC082Cその1(N int, A []int) int {
	bucket := make(map[int]int)

	for _, a := range A {
		bucket[a]++
	}

	var removeCount int

	for num, frequency := range bucket {
		gap := frequency - num
		is供給過多 := gap > 0
		is供給不足 := gap < 0

		if is供給過多 {
			// 過剰な場合は、過剰な分だけを取り除けばよい
			removeCount += gap
		}

		if is供給不足 {
			// 不足の場合は、その数字すべてを消し去るしかない
			removeCount += frequency
		}
	}

	return removeCount
}

func TestAnswerABC082Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want int
	}{
		{"入力例1 すべて必要以上なので削除するだけ", 4, []int{3, 3, 3, 3}, 1},
		{"入力例2 不足分がある場合", 5, []int{2, 4, 1, 4, 2}, 2},
		{"入力例3 なにもしなくていい", 6, []int{1, 2, 2, 3, 3, 3}, 0},
		{"入力例4 空集合が答え", 1, []int{1000000000}, 1},
		{"入力例5", 8, []int{2, 7, 1, 8, 2, 8, 1, 8}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC082Cその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
