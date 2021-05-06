package topic_bit_brute_force

import (
	"testing"
)

// [ABC190C - Bowls and Dishes](https://atcoder.jp/contests/abc190/tasks/abc190_c)
func AnswerABC190Cその1(N, M int, AB [][2]int, K int, CD [][2]int) int {
	// どのボールのどの皿に乗せるかをビット全探索
	patterns := createAllPatterns(N, K, CD)

	maxCount := countOkPatterns(AB, patterns)

	return maxCount
}

func createAllPatterns(N, K int, CD [][2]int) [][]int {
	// 2^K 通りのパターンを格納するスライスを用意する
	var patterns = make([][]int, 1<<K)
	for i := 0; i < 1<<K; i++ {
		patterns[i] = make([]int, N)
	}

	var targetDish int // ボールを乗せる皿番号
	for bit := 0; bit < 1<<K; bit++ {
		for i := 0; i < K; i++ {
			if bit>>i&1 == 1 {
				targetDish = CD[i][0]
			} else {
				targetDish = CD[i][1]
			}
			patterns[bit][targetDish-1]++
		}
	}

	return patterns
}

func countOkPatterns(AB [][2]int, patterns [][]int) int {
	maxCount := 0

	for _, pattern := range patterns {
		count := 0

		for i := 0; i < len(AB); i++ {
			a, b := AB[i][0], AB[i][1]
			if pattern[a-1] > 0 && pattern[b-1] > 0 {
				count++
			}
		}

		maxCount = max(maxCount, count)
	}

	return maxCount
}

func max(x, y int) int {
	if x < y {
		x = y
	}
	return x
}

func TestAnswerABC190Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		M    int
		AB   [][2]int
		K    int
		CD   [][2]int
		want int
	}{
		{"入力例1",
			4, 4,
			[][2]int{
				{1, 2},
				{1, 3},
				{2, 4},
				{3, 4},
			},
			3,
			[][2]int{
				{1, 2},
				{1, 3},
				{2, 3},
			},
			2,
		},

		{"入力例2",
			4, 4,
			[][2]int{
				{1, 2},
				{1, 3},
				{2, 4},
				{3, 4},
			},
			4,
			[][2]int{
				{3, 4},
				{1, 2},
				{2, 4},
				{2, 4},
			},
			4,
		},

		{"入力例3",
			6, 12,
			[][2]int{
				{2, 3},
				{4, 6},
				{1, 2},
				{4, 5},
				{2, 6},
				{1, 5},
				{4, 5},
				{1, 3},
				{1, 2},
				{2, 6},
				{2, 3},
				{2, 5},
			},
			5,
			[][2]int{
				{3, 5},
				{1, 4},
				{2, 6},
				{4, 6},
				{5, 6},
			},
			9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC190Cその1(tt.N, tt.M, tt.AB, tt.K, tt.CD)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
