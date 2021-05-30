package topic_integer_problem

import (
	"reflect"
	"sort"
	"testing"
)

// [ABC180C - Cream puff](https://atcoder.jp/contests/abc180/tasks/abc180_c)
func AnswerABC180Cその1(N int) []int {
	var divisors []int

	for i := 1; i*i <= N; i++ {
		if N%i == 0 {
			divisors = append(divisors, i)

			// N=25, i=5 のとき N/i = 5 となって重複する場合がある
			if N/i != i {
				divisors = append(divisors, N/i)
			}
		}
	}

	sort.Ints(divisors)

	return divisors
}

func TestAnswerABC180Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want []int
	}{
		{"入力例1", 6, []int{1, 2, 3, 6}},
		{"入力例2", 720, []int{1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 15, 16, 18, 20, 24, 30, 36, 40, 45, 48, 60, 72, 80, 90, 120, 144, 180, 240, 360, 720}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC180Cその1(tt.N)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
