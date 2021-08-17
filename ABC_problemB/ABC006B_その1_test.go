package ABC_problemB

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

// [ABC006B - トリボナッチ数列](https://atcoder.jp/contests/abc006/tasks/abc006_2)// [B - おいしいたこ焼きの食べ方](https://atcoder.jp/contests/abc006/tasks/abc006_2)
func AnswerABC006Bその1(N int) int {
	a := 0
	b := 0
	c := 1

	if N == 1 {
		return 0
	}
	if N == 2 {
		return 0
	}
	if N == 3 {
		return 1
	}

	const MOD = 10007
	for i := 3; i < N; i++ {
		a, b, c = b%MOD, c%MOD, (a+b+c)%MOD
	}

	return c
}

func TestAnswerABC006Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{"入力例1", 7, 7},
		{"入力例2", 1, 0},
		{"入力例3", 100000, 7927},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC006Bその1(tt.N)

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Error(diff)
			}
		})
	}
}
