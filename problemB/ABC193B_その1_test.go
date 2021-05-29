package problemB

import (
	"atgoder/lib"
	"math"
	"testing"
)

// [ABC193B - Play Snuke](https://atcoder.jp/contests/abc193/tasks/abc193_b)
func AnswerABC193Bその1(N int, APX [][]int) int {
	minPrice := math.MaxInt32

	for i := 0; i < N; i++ {
		a, p, x := APX[i][0], APX[i][1], APX[i][2]

		canBuy := x-a > 0
		if canBuy {
			minPrice = lib.Min(minPrice, p)
		}
	}

	if minPrice == math.MaxInt32 {
		return -1
	}
	return minPrice
}

func TestAnswerABC193Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		APX  [][]int
		want int
	}{
		{"入力例1",
			3,
			[][]int{
				{3, 9, 5},
				{4, 8, 5},
				{5, 7, 5},
			},
			8},
		{"入力例2",
			3,
			[][]int{
				{5, 9, 5},
				{6, 8, 5},
				{7, 7, 5},
			},
			-1},
		{"入力例3",
			10,
			[][]int{
				{158260522, 877914575, 602436426},
				{24979445, 861648772, 623690081},
				{433933447, 476190629, 262703497},
				{211047202, 971407775, 628894325},
				{731963982, 822804784, 450968417},
				{430302156, 982631932, 161735902},
				{880895728, 923078537, 707723857},
				{189330739, 910286918, 802329211},
				{404539679, 303238506, 317063340},
				{492686568, 773361868, 125660016},
			},
			861648772,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC193Bその1(tt.N, tt.APX)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
