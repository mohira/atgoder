package topic_data_structure

import (
	"reflect"
	"testing"
)

// [ABC137D - Summer Vacation](https://atcoder.jp/contests/abc137/tasks/abc137_d)
func AnswerABC137Dその1(N, M int, AB [][]int) int {
	// 貪欲法: 「真に良い」が確定しないといけない
	return 0
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
