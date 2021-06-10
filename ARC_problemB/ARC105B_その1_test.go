package problemB

import (
	"atgoder/lib"
	"testing"
)

// [ARC105 - MAX-=min](https://atcoder.jp/contests/arc105/tasks/arc105_b)
func AnswerARC105その1(N int, A []int) int {
	// N=2で実験してユークリッドの互除法に気づくのがよいらしい
	// で、引き算を何度やってもGCDは不変ってことがポイントっぽい
	// https://maspypy.com/atcoder-%E5%8F%82%E5%8A%A0%E6%84%9F%E6%83%B3-2020-10-12arc105
	//
	// オレオレ考察メモ
	// ・数直線でイメージすると、最小値付近に集まる感じになる
	// ・減算操作たくさんする → あまりの計算すれば回数減るんじゃね？
	// ・カードに同じ数がいくつあっても結果は変わらない
	//
	// たぶん、最大3回の操作でおわるっぽい
	// 最小値とのあまりを計算していく(あまり0の場合は最小値にする)
	// 3 4 7 8 9 (最小値は3)
	// | | | | |
	// 3 1 1 2 3 (最小値は1)
	// | | | | |
	// 1 1 1 1 1
	//
	// 10 12 16 19 52 900 (最小値は10)
	//  |  |  |  |  |  |
	// 10  2  6  9  2  10 (最小値は2)
	//  |  |  |  |  |  |
	//  2  2  2  1  2  2 (最小値は1)
	//  |  |  |  |  |  |
	//  1  1  1  1  1  1

	min := lib.FindMin(A)
	max := lib.FindMax(A)

	for min != max {
		for i := 0; i < N; i++ {
			if A[i] != min {
				r := A[i] % min
				if r == 0 {
					A[i] = min
				} else {
					A[i] = r
				}
			}
		}

		min = lib.FindMin(A)
		max = lib.FindMax(A)
	}

	return min
}

func TestAnswerARC105その1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want int
	}{
		{
			"入力例1",
			3,
			[]int{2, 6, 6},
			2,
		},
		{"同じ大きさのカードの枚数の影響は受けない",
			10,
			[]int{2, 6, 6, 6, 6, 6, 6, 6, 6, 6},
			2,
		},
		{"",
			3,
			[]int{2, 5, 6},
			1,
		},
		{"",
			5,
			[]int{3, 4, 7, 8, 9},
			1,
		},
		{"",
			5,
			[]int{10, 12, 19, 52, 60},
			1,
		},
		{"入力例2",
			15,
			[]int{546, 3192, 1932, 630, 2100, 4116, 3906, 3234, 1302, 1806, 3528, 3780, 252, 1008, 588},
			42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerARC105その1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
