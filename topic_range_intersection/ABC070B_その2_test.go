package topic_grid

import (
	"testing"
)

// [ABC070B - Two Switches](https://atcoder.jp/contests/abc070/tasks/abc070_b)
func AnswerABC070Bその2(A, B, C, D int) int {
	// 条件を愚直に網羅する場合
	// A<B と C<D は制約条件にあるが、わざわざ書いたほうがわかりやすい
	if A < B && B <= C && C < D {
		return 0
	}
	if A <= C && C <= B && B <= D {
		return B - C
	}
	if A <= C && C < D && C <= B {
		return D - C
	}

	if C < D && D <= A && A < B {
		return 0
	}

	if C <= A && A <= D && D <= B {
		return D - A
	}

	if C <= A && A < B && B <= D {
		return B - A
	}

	panic("条件に漏れがあるよ")
	return 999999999
}

func TestAnswerABC070Bその2(t *testing.T) {
	tests := []struct {
		name string
		A    int
		B    int
		C    int
		D    int
		want int
	}{
		{"入力例1: A<C<B<D 部分重なり", 0, 75, 25, 100, 50},
		{"入力例2: A<B<C<D 重なりナシ", 0, 33, 66, 99, 0},
		{"入力例3: A<C<D<B 包含", 10, 90, 20, 80, 60},

		//
		{"入力例4: C<A<D<B 部分重なり", 25, 100, 0, 75, 50},
		{"入力例5: C<D<A<B 重なりナシ", 66, 99, 0, 33, 0},
		{"入力例6: C<A<B<D 包含", 20, 80, 10, 90, 60},

		// 境界値
		{"入力例7", 0, 75, 0, 75, 75},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC070Bその1(tt.A, tt.B, tt.C, tt.D)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
