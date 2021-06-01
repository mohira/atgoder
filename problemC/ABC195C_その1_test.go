package problemC

import (
	"math"
	"strconv"
	"testing"
)

// [ABC195C - Comma](https://atcoder.jp/contests/abc195/tasks/abc195_c)
func AnswerABC195Cその1(N int) int {
	// 方針その1: 愚直にカウントしていく
	// https://youtu.be/k2fzO_U7RnA?t=1731 あたりの解説と同じ
	// 実装がちょっとややこしめ

	// 桁ごとのコンマの数合計を管理する
	bucket := make(map[int]int)

	for i := 0; i <= 15; i++ {
		n := int(math.Pow10(i) - math.Pow10(i-1)) // 区間にいくつの数字があるか
		comma := (i - 1) / 3                      // コンマの数
		bucket[i] = n * comma                     // その桁のすべての数字の合計コンマ数
	}

	d := len(strconv.Itoa(N)) // 桁数を調べる

	count1 := 0 // Nの桁未満のコンマの合計数
	for keta, commaCount := range bucket {
		if keta < d {
			count1 += commaCount
		}
	}

	n := N - int(math.Pow10(d-1)) + 1 // Nの桁において、Nまでにいくつの整数があるか
	count2 := n * ((d - 1) / 3)       // Nの桁における、Nまでの合計コンマの数

	return count1 + count2
}

func TestAnswerABC195Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{"入力例1", 1_010, 0 + 11},
		{"入力例2", 27182818284590, 107730272137364},

		{"", 10_001, 9000 + 2},
		{"", 100_000, 9000 + 90000 + 1*1},
		{"", 100_001, 9000 + 90000 + 1*2},
		{"", 999_999, 9000 + 90000 + 900000},
		{"", 1_000_000, 9000 + 90000 + 900000 + 2*1},
		{"", 10_000_000, 9000 + 90000 + 900000 + 2*9000000 + 2*1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC195Cその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
