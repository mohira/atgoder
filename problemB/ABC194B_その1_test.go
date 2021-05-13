package problemB

import (
	"math"
	"testing"
)

// [ABC194B - Job Assignment](https://atcoder.jp/contests/abc194/tasks/abc194_b)
func AnswerABC194Bその1(N int, AB [][]int) int {
	// i≠j のときの最小時間
	var ans1 = math.MaxInt32
	for i := 0; i < N; i++ {
		a := AB[i][0]
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			b := AB[j][1]
			ans1 = min(ans1, max(a, b))
		}
	}

	// i=j のときの最小時間
	var ans2 = math.MaxInt32
	for i := 0; i < N; i++ {
		a, b := AB[i][0], AB[i][1]

		ans2 = min(ans2, a+b)
	}

	return min(ans1, ans2)
}

func TestAnswerABC194Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		AB   [][]int
		want int
	}{
		{"入力例1",
			3,
			[][]int{
				{8, 5},
				{4, 4},
				{7, 9},
			},
			5},
		{"入力例2",
			3,
			[][]int{
				{11, 7},
				{3, 2},
				{6, 7},
			},
			5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC194Bその1(tt.N, tt.AB)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
