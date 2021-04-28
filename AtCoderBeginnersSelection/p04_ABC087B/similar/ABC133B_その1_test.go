package similar

import (
	"math"
	"testing"
)

// [ABC133B - Good Distance](https://atcoder.jp/contests/abc133/tasks/abc133_b)
func AnswerABC133Bその1(N int, D int, X [][]int) int {
	var cnt int

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			d := distance(X[i], X[j])
			if isInteger(d) {
				cnt++
			}
		}
	}

	return cnt
}

func isInteger(d float64) bool {
	//return d-math.Floor(d) == 0 // これでもOKっぽい
	return math.Floor(d) == d
}

func distance(y []int, z []int) float64 {
	var sum int

	for i := 0; i < len(y); i++ {
		sum += (y[i] - z[i]) * (y[i] - z[i])
	}

	return math.Sqrt(float64(sum))
}

func TestAnswerABC133Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		D    int
		X    [][]int
		want int
	}{
		{"入力例1",
			3, 2,
			[][]int{
				{1, 2},
				{5, 5},
				{-2, 8}},
			1},
		{"入力例2",
			3, 4,
			[][]int{
				{-3, 7, 8, 2},
				{-12, 1, 10, 2},
				{-2, 8, 9, 3},
			},
			2},
		{"入力例3",
			5, 1,
			[][]int{
				{1},
				{2},
				{3},
				{4},
				{5},
			},
			10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC133Bその1(tt.N, tt.D, tt.X)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}

		})

	}
}
