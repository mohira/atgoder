package problemC

import (
	"atgoder/lib"
	"testing"
)

// [ABC161C -  Replacing Integer](https://atcoder.jp/contests/abc161/tasks/abc161_c)
func AnswerABC161Cその1(N, K int) int {
	// N-Kの正負の切り替えポイントだけ見る
	// N>=K とりあえず引けるところまで引く → Kで割ったあまり
	// N<K  最後にもう一度引くかどうか
	// ex: x=100, k=7
	// 100->93->86->79->...->16->9->2->5->2->5->...
	//                              |
	//                    x>=kの世界 | ループする
	a := N % K
	b := lib.AbsInt(N%K - K)

	return lib.Min(a, b)
}

func TestAnswerABC161Cその1(t *testing.T) {
	tests := []struct {
		name string
		N, K int
		want int
	}{
		{"入力例1", 7, 4, 1},
		{"入力例2", 2, 6, 2},
		{"入力例3", 1000000000000000000, 1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC161Cその1(tt.N, tt.K)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
