package topic_bit_brute_force

import (
	"atgoder/lib"
	"math"
	"testing"
)

// [ABC167C - Skill Up](https://atcoder.jp/contests/abc167/tasks/abc167_c)
func AnswerABC167Cその2(N, M, X int, CA [][]int) int {
	var minCost = math.MaxInt64

	for bit := 0; bit < (1 << N); bit++ {
		skills := make([]int, M)
		cost := 0

		for i := 0; i < N; i++ {
			c := CA[i][0]
			points := CA[i][1:]
			if (bit>>i)&1 == 1 {
				//fmt.Fprintf(os.Stderr, "書籍%d購入 c=%d a=%v\n", i, c, points)
				cost += c

				for j, p := range points {
					skills[j] += p
				}
			}
		}

		// 条件チェック
		if okSkills(skills, X) {
			minCost = lib.Min(minCost, cost)
		}
	}

	if minCost == math.MaxInt64 {
		return -1
	} else {
		return minCost
	}
}

func okSkills(skills []int, X int) bool {
	for _, skill := range skills {
		if skill < X {
			return false
		}
	}

	return true
}

func TestAnswerABC167Cその2(t *testing.T) {
	tests := []struct {
		name string
		N    int
		M    int
		X    int
		CA   [][]int
		want int
	}{
		{"入力例1", 3, 3, 10,
			[][]int{
				{60, 2, 2, 4},
				{70, 8, 7, 9},
				{50, 2, 3, 9},
			},
			120},
		{"入力例2", 3, 3, 10,
			[][]int{
				{100, 3, 1, 4},
				{100, 1, 5, 9},
				{100, 2, 6, 5},
			},
			-1},
		{"入力例3", 8, 5, 22,
			[][]int{
				{100, 3, 7, 5, 3, 1},
				{164, 4, 5, 2, 7, 8},
				{334, 7, 2, 7, 2, 9},
				{234, 4, 7, 2, 8, 2},
				{541, 5, 4, 3, 3, 6},
				{235, 4, 8, 6, 9, 7},
				{394, 3, 6, 1, 6, 2},
				{872, 8, 4, 3, 7, 2},
			},
			1067},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC167Cその2(tt.N, tt.M, tt.X, tt.CA)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
