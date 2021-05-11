package topic_bit_brute_force

import (
	"testing"
)

// [ABC199C - IPFL](https://atcoder.jp/contests/abc199/tasks/abc199_c)
func AnswerABC199Cその2(N int, S string, Q int, TAB [][]int) string {
	//        1 2 | 3 4
	// 反転後: F L | I P
	// 反転前: I P | F L
	//
	//       反転            →    反転前
	// 1文字目(前半; 1<=N  )  → 3文字目(= 1 + N)
	// 3文字目(後半;    N<3)  → 1文字目(= 3 - N)
	//
	// 一般化
	// x文字目について
	//		(x<=N  ) => x = x + N
	//		(   N<x) => x = x - N
	//
	// ゼロインデックスする必要があることに注意

	runes := []rune(S)
	flipped := false

	for _, tab := range TAB {
		t, a, b := tab[0], tab[1], tab[2]

		if t == 2 {
			flipped = !flipped
		} else {
			if flipped {
				if a <= N {
					a += N
				} else {
					a -= N
				}

				if b <= N {
					b += N
				} else {
					b -= N
				}
			}
			a--
			b--
			runes[a], runes[b] = runes[b], runes[a]
		}
	}

	if flipped {
		return string(runes[N:]) + string(runes[:N])
	} else {
		return string(runes)
	}
}

func TestAnswerABC199Cその2(t *testing.T) {
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
			got := AnswerABC199Cその2(tt.N, tt.S, tt.Q, tt.TAB)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
