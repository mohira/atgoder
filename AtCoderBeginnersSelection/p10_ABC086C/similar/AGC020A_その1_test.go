package p10_ABC086C

import (
	"testing"
)

// [AGC020A - Move and Win](https://atcoder.jp/contests/agc010/tasks/agc010_a)
func AnswerAGC020Aその1(N, A, B int) string {
	// A<B なので、前に進むしかない勝ちはない(位置の逆転は起こらない)。
	// 		Alice は 右に進まないと勝てない
	//		Bob   は 左に進まないと勝てない
	// Aliceの手番のときに
	//     Bobとの距離が1マス => 勝てる
	//     Bobとの距離が2マス => 負ける
	//    			 i) 右に進む => Bobの手番のときに1マス差になる => 負ける
	//    			ii) 左に進む => Bobの手番のときに3マス差にできる(瞬間的な延命)
	//   							=> Bobが左に迫ってくる(2マス差)
	//   							=> Aliceが右に進むと1マス差になるので負け確定
	//   							=> Aliceは左に逃げる(2マス差)
	//   							=> 繰り返していくと、Aliceは左端に詰められて、Aliceは右に進むしかなくなる
	//   							=> Bobの手番のときに1マス差になる => 負ける
	//     Bobとの距離が3マス => Bobの手番のときに2マス差にできる => 勝てる
	//     Bobとの距離が4マス
	//    			 i) 右に進める => Bobの手番のときに3マス差になる => 負ける
	//    			ii) 左に進める => Bobの手番のときに5マス差にできる
	//   							=> Bobが左に迫ってくる
	//   							=> 繰り返していくと、左端に詰められて、Aliceは右に進むしかなくなる
	//   							=> Bobの手番のときに1マス差になる => 負ける
	// 的な感じで、マス差の偶奇でたぶんいける
	diff := (B - A) - 1

	if diff%2 == 1 {
		return "Alice"
	} else {
		return "Borys"
	}

}

func TestAnswerAGC020Aその1(t *testing.T) {
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
			got := AnswerAGC020Aその1(tt.N, tt.A, tt.B)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
