package topic_cumlative_sum

import (
	"math"
	"testing"
)

// [ABC098C - Attention](https://atcoder.jp/contests/abc098/tasks/arc098_a)
func AnswerABC098Cその2(N int, S string) int {
	// 西向きの人の累積和がわかれば、東向きの人の人数もわかる(累積和が1つでよいということ)
	// (東向きの人数) = (総数) - (西向きの人数)
	westCumSum := make([]int, 1, N+1) // 「西向きの人」の累積和

	for i, s := range S {
		if s == 'W' {
			westCumSum = append(westCumSum, westCumSum[i]+1)
		} else {
			westCumSum = append(westCumSum, westCumSum[i])
		}
	}

	var minRequiredTurn = math.MaxInt32
	for i := 0; i < N; i++ {
		// リーダーiより西にいて、西側を向いている人(東を向かせたい人の総数)
		w := westCumSum[i]

		// リーダーiより東にいて、東側を向いている人(西を向かせたい人の総数)
		a := N - 1 - i                       // リーダーより東にいる人の総数
		b := westCumSum[N] - westCumSum[i+1] // リーダーより東にいる人で、西を向いている人
		e := a - b

		minRequiredTurn = min(minRequiredTurn, w+e)
	}

	return minRequiredTurn
}

func TestAnswerABC098Cその2(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    string
		want int
	}{
		{"入力例1", 5, "WEEWW", 1},
		{"入力例2", 12, "WEWEWEEEWWWE", 4},
		{"入力例3", 8, "WWWWWEEE", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC098Cその2(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
