package topic_counter

import (
	"reflect"
	"testing"
)

func AnswerABC210Cその1(N, K int, C []int) int {
	return 0
}

func TestABC210Cその1(t *testing.T) {
	tests := []struct {
		name string
		N, K int
		C    []int
		want int
	}{
		{
			"入力例1",
			7, 3,
			[]int{1, 2, 1, 2, 3, 3, 1},
			3,
		},
		{
			"入力例2",
			5, 5,
			[]int{4, 4, 4, 4, 4},
			1,
		},
		{
			"入力例3",
			10, 6,
			[]int{1, 2, 1, 2, 3, 1, 4, 1, 1, 4},
			4,
		},

		// 自作入力例
		{"", 7, 6, []int{3, 1, 4, 1, 5, 9, 2}, 5}, // またぐと嬉しいやつ
		{"", 5, 1, []int{1, 3, 1, 1, 4}, 1},
		{"", 5, 2, []int{1, 3, 1, 1, 4}, 2},
		{"", 5, 3, []int{1, 3, 1, 1, 4}, 2}, // またげないやつ
		{"", 5, 4, []int{1, 3, 1, 1, 4}, 3},
		{"", 5, 5, []int{1, 3, 1, 1, 4}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC210Cその1(tt.N, tt.K, tt.C)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
