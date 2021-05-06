package topic_bit_brute_force

import (
	"testing"
)

// [ABC128C - Switches](https://atcoder.jp/contests/abc128/tasks/abc128_c)
func AnswerABC128Cその1(N, M int, KS [][]int, P []int) int {
	var 全灯となるスイッチのパターン数 = 0

	allSwitchesPatterns := createAllSwitchesPatterns(N)

	for _, switchesPattern := range allSwitchesPatterns {
		if あるパターンで全灯する(M, KS, P, switchesPattern) {
			全灯となるスイッチのパターン数++
		}
	}

	return 全灯となるスイッチのパターン数
}

func createAllSwitchesPatterns(N int) [][]int {
	var allPatterns = make([][]int, 1<<N)

	for bit := 0; bit < (1 << N); bit++ {
		switchesPattern := make([]int, N)

		for i := 0; i < N; i++ {
			if (bit>>i)&1 == 1 {
				// スイッチi を押すとき
				switchesPattern[i] = 1
			}
		}
		allPatterns[bit] = switchesPattern
	}

	return allPatterns
}

func あるパターンで全灯する(M int, KS [][]int, P []int, switchesPattern []int) bool {
	for i := 0; i < M; i++ {
		if !ある電球が点灯する(switchesPattern, KS[i][1:], P[i]) {
			return false
		}
	}

	return true
}

func ある電球が点灯する(switchPattern []int, S []int, p int) bool {
	onSwitchesCount := 0
	for i := 0; i < len(S); i++ {
		if switchPattern[S[i]-1] == 1 {
			onSwitchesCount++
		}
	}

	if onSwitchesCount%2 == p {
		return true
	} else {
		return false
	}
}

func TestAnswerABC128Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		M    int
		KS   [][]int
		P    []int
		want int
	}{
		{"入力例1",
			2, 2,
			[][]int{
				{2, 1, 2},
				{1, 2},
				{0, 1},
			},
			[]int{0, 1},
			1,
		},
		{"入力例2",
			2, 3,
			[][]int{
				{2, 1, 2},
				{1, 1},
				{1, 2},
			},
			[]int{0, 0, 1},
			0,
		},
		{"入力例3",
			5, 2,
			[][]int{
				{3, 1, 2, 5},
				{2, 2, 3},
				{1, 0},
			},
			[]int{1, 0},
			8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC128Cその1(tt.N, tt.M, tt.KS, tt.P)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
