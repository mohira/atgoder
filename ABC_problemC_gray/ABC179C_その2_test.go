package ABC_problemC_gray

import (
	"testing"
)

// [ABC179C - A x B + C](https://atcoder.jp/contests/abc179/tasks/abc179_c)
func AnswerABC179Cその2(N int) int {
	// A,Bが決まると、Cは決まる
	// Aを固定するとBも決まるので、実はBの探索は不要
	// まずは、Aを固定する
	// AB + C = N
	//      C = N - AB
	// で、Cは正の整数なので
	//      C = N - AB >= 1
	// つまり、
	//          N - AB >= 1 を満たすA,Bの組み合わせ を求めればよい
	// ABに着目して言い換えると、
	// 		AB <= N-1
	// また、条件を満たすBは、(N-1)/A [個] だけある
	// よって、Aを固定して、(N-1)/A [個]を加算すればOK

	var count int
	for A := 1; A <= N-1; A++ {
		// AB=N-1 を満たすBの個数
		count += (N - 1) / A
	}
	return count
}

func TestAnswerABC179Cその2(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{"入力例1", 3, 3},
		{"入力例2", 100, 473},
		{"入力例3", 1000000, 13969985},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC179Cその2(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
