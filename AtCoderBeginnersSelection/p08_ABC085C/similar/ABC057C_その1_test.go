package p08_ABC085C

import (
	"math"
	"testing"
)

// [ABC057C - Digits in Multiplication](https://atcoder.jp/contests/abc057/tasks/arc086_a)
func AnswerABC057Cその1(N int) int {
	// 探索範囲、1 <= a <= 10^10 は間に合わない

	// N = ab = ba なので、
	// で、aが大きくなると、bは小さくなる
	// aが最大となるのは、a = √N のときで、このとき、b = √N で bも最大
	// なので、 1 <= a <= √N だけ調べればいいっぽい
	// 一方、 a > √N だとすると b の値が決められないと思うから(たぶん)
	// https://drken1215.hatenablog.com/entry/2019/09/16/221500
	//
	//
	// 愚直に調べてみる
	// N=8なら、
	// a | 1 2 3 | 4 5 6 7 8
	// b | 8 4 x | 2 x x x 1
	//             ↑
	//             すでに調べている！(2,4)と(4,2)は同じ
	// N=16なら、
	// a |  1 2 3 4 | 5 6 7 8 9 10 11 12 13 14 15 16
	// b | 16 8 x 4 | x x x 2 x  x  x  x  x  x  x  1
	//              ↑
	//            ここを境に、以降はすでに調べている！(2,4)と(4,2)は同じ
	var minF int = 1e+10

	maxA := int(math.Sqrt(float64(N)))
	for a := 1; a <= maxA; a++ {
		b := findB(N, a)

		if b > 0 {
			f := F(a, b)
			if f < minF {
				minF = f
			}
		}
	}
	return minF
}

func digitNum(n int) int {
	// 常用対数を使っても桁数は取れる
	return int(math.Log10(float64(n))) + 1

	// 素直にやるパターン
	//var digit int
	//
	//for n > 0 {
	//	digit++
	//
	//	n /= 10
	//}
	//
	//return digit
}

func F(a int, b int) int {
	digitA := digitNum(a)
	digitB := digitNum(b)

	if digitA < digitB {
		return digitB
	} else {
		return digitA
	}
}

func findB(n int, a int) int {
	if n%a == 0 {
		return n / a
	} else {
		return 0
	}

}

func TestAnswerABC057Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{"入力例1", 10000, 3},
		{"入力例2", 1000003, 7},
		{"入力例3", 9876543210, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC057Cその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
