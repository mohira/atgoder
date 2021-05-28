package problemB

import (
	"testing"
)

// [ABC033B - 町の合併](https://atcoder.jp/contests/abc033/tasks/abc033_b)
func AnswerABC033Bその1(N int, S []string, P []int) string {
	totalPeople := sumInts(P)

	for i := 0; i < N; i++ {
		s, p := S[i], P[i]

		if totalPeople < p*2 {
			return s
		}
	}
	return "atcoder"
}

func TestAnswerABC033Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		S    []string
		P    []int
		want string
	}{
		{
			"入力例1",
			4,
			[]string{"unagi", "usagi", "snuke", "smeke"},
			[]int{20, 13, 42, 7},
			"snuke",
		},
		{
			"入力例2",
			5,
			[]string{"a", "b", "c", "d", "e"},
			[]int{10, 20, 30, 40, 100},
			"atcoder",
		},
		{
			"入力例3",
			5,
			[]string{"a", "b", "c", "d", "e"},
			[]int{10, 20, 30, 40, 100},
			"atcoder",
		},
		{
			"整数の /2 演算だと過半数にならないやつ",
			3,
			[]string{"a", "b", "c"},
			[]int{2, 3, 2},
			"atcoder",
		},
		{
			"入力例3",
			14,
			[]string{"yasuzuka", "uragawara", "oshima", "maki", "kakizaki", "ogata", "kubiki", "yoshikawa", "joetsu", "nakago", "itakura", "kiyosato", "sanwa", "nadachi"},
			[]int{3340, 4032, 2249, 2614, 11484, 10401, 9746, 5142, 100000, 4733, 7517, 3152, 6190, 3169},
			"joetsu",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC033Bその1(tt.N, tt.S, tt.P)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
