package topic_next_permutation

import (
	"math"
	"testing"
)

// [ABC145C - Average Length](https://atcoder.jp/contests/abc145/tasks/abc145_c)
func AnswerABC145Bその2(N int, XY [][2]int) float64 {
	// Permutatorより明快だけど、遅い
	a := make([]int, 0, N)
	for i := 0; i <= N-1; i++ {
		a = append(a, i)
	}

	var sum float64

	for _, p := range permutation(a) {
		for i := 0; i < len(p)-1; i++ {
			p1 := XY[p[i]]
			p2 := XY[p[i+1]]

			sum += distance(p1, p2)
		}
	}

	return sum / float64(len(permutation(a)))
}

func TestAnswerABC145Bその2(t *testing.T) {
	tests := []struct {
		name string
		N    int
		XY   [][2]int
		want float64
	}{
		{"入力例1",
			3,
			[][2]int{
				{0, 0},
				{1, 0},
				{0, 1},
			}, 2.2761423749},

		{"入力例2",
			2,
			[][2]int{
				{-879, 981},
				{-866, 890},
			}, 91.9238815543},

		{"入力例3",
			8,
			[][2]int{
				{-406, 10},
				{512, 859},
				{494, 362},
				{-955, -475},
				{128, 553},
				{-986, -885},
				{763, 77},
				{449, 310},
			}, 7641.9817824387},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC145Bその2(tt.N, tt.XY)

			if math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
