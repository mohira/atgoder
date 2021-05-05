package topic_cumlative_sum

import (
	"testing"
)

// [ABC087C - Candies](https://atcoder.jp/contests/abc087/tasks/arc090_a)
func AnswerABC087Cその2(N int, A [2][]int) int {
	// 累積和でやると楽
	//    1 3 5 7  9
	//    2 4 6 8 10
	//
	// 上段(A[0]) と 下段(A[1]) それぞれの累積和
	//
	// i: 0  1  2  3  4  5
	// 0: 0  1  4  9 16 25
	// 1: 0  2  6 10 18 28
	//
	// i=3 で下に移動したとすると
	// 上段での獲得アメ数 = 4 (i=3を読めばいいだけ)
	// 下段での獲得アメ数 = 28 - 6 = (下段のアメ総数) - (i=2(3-1)で獲得できたはずのアメ数)

	// 累積和をもとめる
	line1 := getCumulativeSum(N, A[0])
	line2 := getCumulativeSum(N, A[1])

	var maxCandies int

	for i := 1; i <= N; i++ {
		// iのタイミングで下移動をしたとき
		candiesFromLine1 := line1[i]
		candiesFromLine2 := line2[N] - line2[i-1]

		totalCandies := candiesFromLine1 + candiesFromLine2

		if maxCandies < totalCandies {
			maxCandies = totalCandies
		}
	}

	return maxCandies
}

func getCumulativeSum(N int, A []int) []int {
	sums := make([]int, 0, N+1)
	sums = append(sums, 0)

	var s int
	for i := 0; i < N; i++ {
		s += A[i]
		sums = append(sums, s)
	}

	return sums
}

func TestAnswerABC087Cその2(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    [2][]int
		want int
	}{
		{"わかりやすいやつ", 3,
			[2][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			16,
		},
		{"入力例1", 5,
			[2][]int{
				{3, 2, 2, 4, 1},
				{1, 2, 2, 2, 1},
			},
			14,
		},

		{"入力例2", 4,
			[2][]int{
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
			},
			5,
		},

		{"入力例3", 7,
			[2][]int{
				{3, 3, 4, 5, 4, 5, 3},
				{5, 3, 4, 4, 2, 3, 2},
			},
			29,
		},

		{"入力例4", 1,
			[2][]int{
				{2},
				{3},
			},
			5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC087Cその2(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
