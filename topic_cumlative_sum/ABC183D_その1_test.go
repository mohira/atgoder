package topic_cumlative_sum

import (
	"testing"
)

// [ABC183D - Water Heater](https://atcoder.jp/contests/abc183/tasks/abc183_d)
func AnswerABC183Dその1(N, W int, STP [][]int) string {
	// いもす法
	const nMax = 200010
	imos := make([]int, nMax)
	for i := 0; i < N; i++ {
		s, t, p := STP[i][0], STP[i][1], STP[i][2]
		t--
		imos[s] += p
		imos[t+1] -= p
	}

	// わざわざ累積和のためのスライスを用意しなくてもいい
	for i := 0; i < nMax-1; i++ {
		imos[i+1] += imos[i]

		if W < imos[i] {
			return "No"
		}
	}

	return "Yes"
}

func TestAnswerABC183Dその1(t *testing.T) {
	tests := []struct {
		name string
		N, W int
		STP  [][]int
		want string
	}{
		{"入力例1",
			4, 10,
			[][]int{
				{1, 3, 5},
				{2, 4, 4},
				{3, 10, 6},
				{2, 4, 1},
			},
			"No"},
		{"入力例2",
			4, 10,
			[][]int{
				{1, 3, 5},
				{2, 4, 4},
				{3, 10, 6},
				{2, 3, 1},
			},
			"Yes"},

		{"入力例3",
			6, 1000000000,
			[][]int{
				{0, 200000, 999999999},
				{2, 20, 1},
				{20, 200, 1},
				{200, 2000, 1},
				{2000, 20000, 1},
				{20000, 200000, 1},
			},
			"Yes"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC183Dその1(tt.N, tt.W, tt.STP)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
