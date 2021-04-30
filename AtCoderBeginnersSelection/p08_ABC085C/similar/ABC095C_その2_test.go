package p08_ABC085C

import (
	"math"
	"testing"
)

// [ABC095C - Half and Half](https://atcoder.jp/contests/abc095/tasks/arc096_a)
func AnswerABC095Cその2(A, B, C, X, Y int) int {
	// 発想MEMO: 既知としたときに絞り込みに大きく影響する変数は何か？

	// nA,nB,nCをそれぞれのピザの購入枚数とすると
	//
	// 金額 = A*nA + B*nB + C*2*nC
	//
	// ややこしいので、あらためて C = 2C とすると次のようにできて考えやすい。
	//
	// 金額 = A*nA + B*nB + C*nC
	//
	// 今回、満たすべきピザ枚数条件は、こうなる。
	//
	// 1) X <= nA + nC
	// 2) Y <= nA + nC
	//
	// 上記 1) および 2) を満たした上での、金額の最小値を求めたい。
	//
	// ここで、最小値を求めたいので、等号が成り立つ場合だけを考えればよい。
	//
	// つまり、こう。
	//
	// 3) X = nA + nC
	// 4) Y = nA + nC
	//
	// さらに、nA, nB, nCには非負条件があるので
	//
	// 5) X = (nA + nC) >= 0 ⇔ nA = (X - nC) >= 0
	// 6) Y = (nB + nC) >= 0 ⇔ nB = (Y - nC) >= 0
	//
	// ここでポイントとなるのは、「nCが決まれば、nAとnBも決まる」ということ(X,Yは既知)
	//
	// そこで、nCの取りうる範囲を考えると、こうなる。
	//
	// 0 <= nC <= max(X, Y) <= 10^5
	//
	//
	// あとは、この探索範囲において、先の条件 5)と6)を満足する最小の金額をみつければOK。

	C = 2 * C

	var minCost = math.MaxInt32
	for nC := 0; nC <= max(X, Y); nC++ {
		// 非負条件の確認
		// Cピザで枚数が足りているなら、AピザとBピザは買わない(n=0)という選択もあり得る
		nA := max(X-nC, 0)
		nB := max(Y-nC, 0)

		cost := A*nA + B*nB + C*nC
		if cost < minCost {
			minCost = cost
		}
	}

	return minCost
}

func max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func TestAnswerABC096Cその2(t *testing.T) {
	tests := []struct {
		name string
		A    int
		B    int
		C    int
		X    int
		Y    int
		want int
	}{
		{"入力例1", 1500, 2000, 1600, 3, 2, 7900},
		{"入力例2", 1500, 2000, 1900, 3, 2, 8500},
		{"入力例3", 1500, 2000, 500, 90000, 100000, 100000000},
		{"入力例1", 1500, 2000, 1750, 3, 2, 8500},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC095Cその2(tt.A, tt.B, tt.C, tt.X, tt.Y)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
