package ABC_problemC_gray

import (
	"atgoder/lib"
	"strings"
	"testing"
)

// [ABC120C - Unification](https://atcoder.jp/contests/abc120/tasks/abc120_c)
func AnswerABC120Cその1(S string) int {
	// 0と1が1個以上ある場合は、必ず隣り合っている ⇔ 消せる
	// 処理を続けると最終的には、次の3パターン
	//   i) 0のみ
	//  ii) 1のみ
	// iii) 全部消える
	count0 := strings.Count(S, "0")
	count1 := strings.Count(S, "1")

	return lib.Min(count0, count1) * 2
}

func TestAnswerABC120Cその1(t *testing.T) {
	tests := []struct {
		name string
		S    string
		want int
	}{
		{"入力例1", "0011", 4},
		{"入力例2", "11011010001011", 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC120Cその1(tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
