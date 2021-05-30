package topic_amari

import (
	"atgoder/lib"
	"testing"
)

// [ABC140C - Maximal Value](https://atcoder.jp/contests/abc140/tasks/abc140_c)
func AnswerABC140Cその1(N int, B []int) int {
	var A = make([]int, N)

	// Aの末項はBの末項となることが確定
	A[N-1] = B[N-2]

	// 後ろから決めていく(Aの第2項まで); Bの情報だけで決まる
	for i := len(B) - 1; i > 0; i-- {
		A[i] = lib.Min(B[i], B[i-1])
	}

	// Aの初項はなるべく大きいほうがいい
	A[0] = lib.Max(B[0], A[1])

	return lib.SumInts(A)
}

func TestAnswerABC140Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		B    []int
		want int
	}{
		{"入力例1", 3, []int{2, 5}, 9},
		{"入力例2", 2, []int{3}, 6},
		{"入力例3", 6, []int{0, 153, 10, 10, 23}, 53},
		{"WA", 6, []int{1, 0, 1, 0, 1}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC140Cその1(tt.N, tt.B)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
