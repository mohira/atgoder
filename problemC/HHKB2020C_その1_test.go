package problemC

import (
	"fmt"
	"strings"
	"testing"
)

// [HHKB2020C - Neq Min](https://atcoder.jp/contests/hhkb2020/tasks/hhkb2020_c)
func AnswerHHKB2020Cその1(N int, P []int) string {
	// すでに登場するか？ を確認する方法だとO(N^2)で詰む
	// 考察ポイント
	// ・連想配列じゃなくて整数バケットでの出力をみると傾向が見えやすい
	// ・最小値は徐々に大きくなっていく傾向がある(前回の最小値以上になる)
	// ・idxを徐々にずらしていくことが期待できる
	// ・未登場かどうかチェックではなく、現在の最小値のidxを保持していく)
	var s = make([]string, N)
	bucket := make([]int, 200001)

	min := 0

	for i, p := range P {
		bucket[p]++

		for {
			if bucket[min] == 0 {
				s[i] = fmt.Sprintf("%d", min)
				break
			}
			min++
		}
	}

	return strings.Join(s, "\n")
}

func TestAnswerHHKB2020Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		P    []int
		want string
	}{
		{"入力例1", 4, []int{1, 1, 0, 2}, "0\n0\n2\n3"},
		{"入力例2", 10, []int{5, 4, 3, 2, 1, 0, 7, 7, 6, 6}, "0\n0\n0\n0\n0\n6\n6\n6\n8\n8"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerHHKB2020Cその1(tt.N, tt.P)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
