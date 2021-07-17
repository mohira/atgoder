package C

import (
	"reflect"
	"testing"
)

func AnswerC(N, K int, C []int) int {
	return 0
}

func TestC(t *testing.T) {
	tests := []struct {
		name string
		N, K int
		C    []int
		want int
	}{
		{
			"入力例1",
			7, 3,
			[]int{0, 0, 0},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerC(tt.N, tt.K, tt.C)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
