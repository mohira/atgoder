package p10_ABC086C

import (
	"testing"
)

// [AGC010A - Addition](https://atcoder.jp/contests/agc010/tasks/agc010_a)
func AnswerAGC010Aその2(N int, A []int) string {
	// もし、1つの値にできた場合
	// 		残る数Nは、 N = sum(A) となる(足し算を繰り返しているだけだからね)
	//
	// 		で、1つの値になっているということは、必ず偶数になるはず。
	// 		なぜなら、「偶奇が同じ値を合計する操作」をした結果は、必ず偶数だから
	//
	//  	結局、すべての数字の合計が偶数なら1つだけ残せるってこと
	sum := calcSum(A)
	if sum%2 == 0 {
		return "YES"
	} else {
		return "NO"
	}
}

func calcSum(A []int) int {
	var sum int
	for _, a := range A {
		sum += a
	}
	return sum
}

func TestAnswerAGC010Aその2(t *testing.T) {
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
			got := AnswerAGC010Aその2(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
