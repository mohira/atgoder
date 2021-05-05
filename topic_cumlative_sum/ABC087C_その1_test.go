package topic_cumlative_sum

import (
	"testing"
)

// [ABC087C - Candies](https://atcoder.jp/contests/abc087/tasks/arc090_a)
func AnswerABC087Cその1(N int, A [2][]int) int {
	// 移動するパターンは N+1 通り なので パターンを網羅すると O(N)
	// 最大値を求めるのに O(N)
	// よって、O(N^2)

	// 右移動or下移動 なので、基本右に進み、どのタイミングで下移動するかでパターンが決まる
	var max int

	for i := 0; i < N; i++ {
		candies := make([]int, 0, N+1)
		candies = append(candies, A[0][0])

		x, y := 0, 0
		stepCount := 0

		for len(candies) <= N {
			// i回目で下に移動する
			if stepCount == i {
				x++
			} else {
				y++
			}
			candies = append(candies, A[x][y])
			stepCount++
		}
		totalCandies := sum(candies)

		if max < totalCandies {
			max = totalCandies
		}

	}

	return max
}

func sum(A []int) int {
	var s int
	for i := 0; i < len(A); i++ {
		s += A[i]
	}
	return s
}

func TestAnswerABC087Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    [2][]int
		want int
	}{
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
			got := AnswerABC087Cその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
