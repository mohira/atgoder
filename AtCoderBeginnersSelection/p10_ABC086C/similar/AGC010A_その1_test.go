package p10_ABC086C

import (
	"testing"
)

// [AGC010A - Addition](https://atcoder.jp/contests/agc010/tasks/agc010_a)
func AnswerAGC010Aその1(N int, A []int) string {
	// 偶数がN個あるとき、
	// 		Nが偶数 => 1個の偶数 が残る
	// 		Nが奇数 => 1個の偶数 が残る
	//
	// 奇数がM個あるとき、
	//      Mが偶数 => 1個の偶数 が残る
	//      Mが奇数 => 1個の偶数 と 1個の奇数 が残る
	//
	// 表にすると、こんな感じ
	//
	//  N | M
	// 偶 | 偶 => {偶, 偶}     -> {偶}    YES
	// 偶 | 奇 => {偶, 偶, 奇} -> {偶, 奇} NO
	// 奇 | 偶 => {偶, 偶}     -> {偶}    YES
	// 奇 | 奇 => {偶, 偶, 奇} -> {偶, 奇} NO
	nOddNumber := countOddNumber(A)

	奇数要素が奇数個ある := nOddNumber%2 == 1

	if 奇数要素が奇数個ある {
		return "NO"
	}

	return "YES"
}

func countOddNumber(A []int) int {
	var c int

	for _, a := range A {
		if a%2 == 1 {
			c++
		}
	}

	return c
}

func TestAnswerAGC010Aその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want string
	}{
		{"入力例1", 3, []int{1, 2, 3}, "YES"},
		{"入力例2", 5, []int{1, 2, 3, 4, 5}, "NO"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerAGC010Aその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
