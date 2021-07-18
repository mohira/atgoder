package ABC_problemB

import (
	"testing"
)

// [ABC015B - 高橋くんの集計](https://atcoder.jp/contests/abc015/tasks/abc015_2)
func AnswerABC015Bその1(N int, A []int) int {
	total := 0
	count := 0
	for _, a := range A {
		if a > 0 {
			total += a
			count++
		}
	}

	return (total + count - 1) / count

	//p, q := total/count, total%count
	//if q > 0 {
	//	p++
	//}
	//return p
}

func TestAnswerABC015Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want int
	}{
		{"入力例1", 4, []int{0, 1, 3, 8}, 4},
		{"入力例2", 5, []int{1, 4, 9, 10, 15}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC015Bその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
