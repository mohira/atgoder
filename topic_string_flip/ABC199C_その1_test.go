package topic_bit_brute_force

import (
	"strings"
	"testing"
)

// [ABC199C - IPFL](https://atcoder.jp/contests/abc199/tasks/abc199_c)
func AnswerABC199Cその1_TLE(N int, S string, Q int, TAB [][]int) string {
	var strs = make([][]string, 2)
	strs[0] = strings.Split(S[:N], "")
	strs[1] = strings.Split(S[N:], "")

	for _, tab := range TAB {
		t, a, b := tab[0], tab[1], tab[2]

		if t == 2 {
			strs[0], strs[1] = strs[1], strs[0]
		} else {
			if b <= N {
				// 前半だけでの入れ替え
				i := a - 1
				j := b - 1
				strs[0][i], strs[0][j] = strs[0][j], strs[0][i]
			} else if N < a {
				// 後半だけでの入れ替え
				i := a - N - 1
				j := b - N - 1
				strs[1][i], strs[1][j] = strs[1][j], strs[1][i]
			} else {
				// 前半と後半での入れ替え
				i := a - 1
				j := b - N - 1
				strs[0][i], strs[1][j] = strs[1][j], strs[0][i]
			}
		}
	}

	var ans string
	for i := 0; i < 2; i++ {
		for j := 0; j < N; j++ {
			ans += strs[i][j]
		}
	}

	return ans
}

func TestAnswerABC199Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		S    string
		Q    int
		TAB  [][]int
		want string
	}{
		{"入力例1",
			2,
			"FLIP",
			2,
			[][]int{
				{2, 0, 0},
				{1, 1, 4},
			},
			"LPFI",
		},
		{"入力例2",
			2,
			"FLIP",
			6,
			[][]int{
				{1, 1, 3},
				{2, 0, 0},
				{1, 1, 2},
				{1, 2, 3},
				{2, 0, 0},
				{1, 1, 4},
			},
			"ILPF"},
		{"後半の入れ替えがおきるケース",
			2,
			"FLIP",
			1,
			[][]int{
				{1, 3, 4},
			},
			"FLPI",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC199Cその1_TLE(tt.N, tt.S, tt.Q, tt.TAB)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
