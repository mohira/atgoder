package ABC_problemC_gray

import (
	"atgoder/lib"
	"testing"
)

// [ABC188C -  ABC Tournament](https://atcoder.jp/contests/abc188/tasks/abc188_c)
func AnswerABC188Cその2(N int, A []int) int {
	// 素直にシミュレーションして勝者のリストを更新していく
	players := A

	// 決勝戦の対戦カードが決まるまで(==2人残るまで)続ける
	for len(players) > 2 {
		winners := make([]int, 0, len(players)/2)

		for i := 0; i < len(players)-1; i += 2 {
			p1, p2 := players[i], players[i+1]

			winners = append(winners, lib.Max(p1, p2))
		}
		players = winners
	}

	// 準優勝の値は決勝戦の対戦でレートが低い方
	secondValue := lib.Min(players[0], players[1])

	for i, a := range A {
		if a == secondValue {
			return i + 1
		}
	}

	return 0
}

func TestAnswerABC188Cその2(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want int
	}{
		{"入力例1", 2, []int{1, 4, 2, 5}, 2},
		{"入力例2", 2, []int{3, 1, 5, 4}, 1},
		{"入力例3", 4, []int{6, 13, 12, 5, 3, 7, 10, 11, 16, 9, 8, 15, 2, 1, 14, 4}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC188Cその2(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
