package p10_ABC086C

import (
	"testing"
)

// [AGC020A - Move and Win](https://atcoder.jp/contests/agc010/tasks/agc010_a)
func AnswerAGC020Aその2(N, A, B int) string {
	D := B - A

	if D%2 == 0 {
		// Dが偶数
		// Aliceの手番: Aliceのコマは偶数
		// Borysの手番: Borysのコマは偶数

		// Dが偶数 => D!=1 => コマが隣接していない => Aliceは必ずコマを動かせる
		// Aliceは必ずコマを動かせる => いつか、N-1に到達できる => そのときボリスはN(==右端) => Aliceの勝ち
		return "Alice"
	} else {
		// Dが奇数
		// Aliceの手番: Aliceのコマは奇数
		// Borysの手番: Borysのコマは偶数

		// Aliceでの議論と同様で、ボリスはいつか 2ますに到達 => そのときアリスは1マス(==左端) => Borysの勝ち
		return "Borys"
	}
}

func TestAnswerAGC020Aその2(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    int
		B    int
		want string
	}{
		{"入力例1", 5, 2, 4, "Alice"},
		{"入力例2", 2, 1, 2, "Borys"},
		{"入力例3", 58, 23, 42, "Borys"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerAGC020Aその2(tt.N, tt.A, tt.B)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
