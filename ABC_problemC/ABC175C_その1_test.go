package ABC_problemC

import (
	"atgoder/lib"
	"testing"
)

// [ABC175C - Compass Walking](https://atcoder.jp/contests/abc175/tasks/abc175_c)
func AnswerABC175Cその1(X, K, D int) int {
	// 原点からの距離が問題なので、絶対値で扱う(X>=0の世界で処理できる)
	X = lib.AbsInt(X)

	// "残り回数を気にせず"、なるべく原点に近づくとする
	// 原点付近では、行き過ぎがあるので、最終位置候補は2択になる

	// i) 往復パターンに入る前に行動回数がつきる
	if K <= X/D {
		return X - K*D
	}

	// ii) 往復パターンに突入する
	// p を 往復パターンの初期位置とすると、
	// ここからは -D と +D を繰り返す
	// 残り偶数回: 2回で戻ってこれるので最終位置はp
	// 残り奇数回: 最終位置はp-D (答えは距離なのでabsとる)
	p := X % D // 往復パターン初期位置(+位置)
	K -= X / D // pに到達するまでの消費手数を差っ引く

	if K%2 == 0 {
		return p
	} else {
		return lib.AbsInt(p - D)
	}
}

func TestAnswerABC175Cその1(t *testing.T) {
	tests := []struct {
		name    string
		X, K, D int
		want    int
	}{
		{"入力例1", 6, 2, 4, 2},
		{"入力例2", 7, 4, 3, 1},
		{"入力例3", 10, 1, 2, 8},
		{"入力例4", 1000000000000000, 1000000000000000, 1000000000000000, 1000000000000000},

		{"入力例2でマイナス位置が最終地点(残Kが奇数)", 7, 5, 3, 2},
		{"往復パターンに入る前に移動を使い切る", 7, 1, 3, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC175Cその1(tt.X, tt.K, tt.D)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
