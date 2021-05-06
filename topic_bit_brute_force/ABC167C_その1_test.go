package topic_bit_brute_force

import (
	"math"
	"testing"
)

// [ABC167C - Skill Up](https://atcoder.jp/contests/abc167/tasks/abc167_c)
func AnswerABC167Cその1(N, M, X int, CA [][]int) int {
	// 購入する本の部分集合をビット全探索で作る
	// オーダーは、O(N^2 MN)
	var can bool
	var minTotalPrice = math.MaxInt32

	for bit := 0; bit < (1 << N); bit++ {
		totalPrice := 0

		var bookSkills = make([][]int, 0, N)

		for i := 0; i < N; i++ {
			//if bit&(1<<i) > 0 でもOK
			if (bit>>i)&1 == 1 {
				price, skills := CA[i][0], CA[i][1:]

				bookSkills = append(bookSkills, skills)
				totalPrice += price
			}
		}

		understandingLevels := to理解度(bookSkills, M)

		if ok(understandingLevels, X) {
			can = true

			minTotalPrice = min(totalPrice, minTotalPrice)
		}
	}

	if can {
		return minTotalPrice
	} else {
		return -1
	}
}

func min(x, y int) int {
	if x < y {
		y = x
	}
	return y
}

// 各理解度が条件を見対しているかをチェックする
func ok(skills []int, X int) bool {
	for _, skill := range skills {
		if skill < X {
			return false
		}
	}
	return true
}

// 複数の書籍からM個のアルゴリズムそれぞれの理解度の合計をとる
func to理解度(subset [][]int, M int) []int {
	var understandingLevels = make([]int, M)

	for i := 0; i < len(subset); i++ {
		for j := 0; j < M; j++ {
			understandingLevels[j] += subset[i][j]
		}
	}
	return understandingLevels
}

func TestAnswerABC167Cその1(t *testing.T) {
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
			got := AnswerABC167Cその1(tt.N, tt.M, tt.X, tt.CA)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
