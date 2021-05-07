package topic_next_permutation

import (
	"math"
	"testing"
)

// [ABC145C \- Average Length](https://atcoder.jp/contests/abc145/tasks/abc145_c)
func AnswerABC145Bその1(N int, XY [][2]int) float64 {
	a := make([]int, 0, N)
	for i := 0; i <= N-1; i++ {
		a = append(a, i)
	}

	permutations := NewPermutator(IntSlice(a))

	var sum float64

	n := 0 // パターン数
	for permutations.Next() {
		n++
		for i := 0; i < len(a)-1; i++ {
			p1 := XY[a[i]]
			p2 := XY[a[i+1]]

			sum += distance(p1, p2)
		}
	}

	return sum / float64(n)
}

func distance(p1, p2 [2]int) float64 {
	x1, y1 := p1[0], p1[1]
	x2, y2 := p2[0], p2[1]

	x := (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)

	return math.Sqrt(float64(x))
}

func TestAnswerABC145Bその1(t *testing.T) {
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
			got := AnswerABC145Bその1(tt.N, tt.XY)

			if math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
