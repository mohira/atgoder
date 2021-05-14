package problemB

import (
	"testing"
)

// [ABC189B - Alcoholic](https://atcoder.jp/contests/abc189/tasks/abc189_b)
func AnswerABC189Bその1(N, X int, VP [][]int) int {
	var totalAlcohol int

	for i := 0; i < N; i++ {
		v, p := VP[i][0], VP[i][1]

		alcohol := v * p

		totalAlcohol += alcohol

		if X*100 < totalAlcohol {
			return i + 1
		}
	}
	return -1
}

func TestAnswerABC189Bその1(t *testing.T) {
	tests := []struct {
		name string
		N, X int
		VP   [][]int
		want int
	}{
		{"float64の誤差が問題になるケース",
			1, 7,
			[][]int{
				{25, 28},
			},
			-1},

		{"最小のアルコール量で酔うケース",
			1, 0,
			[][]int{
				{1, 1},
			},
			1},

		{"入力例1",
			2, 15,
			[][]int{
				{200, 5},
				{350, 3},
			},
			2},

		{"入力例2 アルコールの摂取量がちょうど X ml のとき、高橋君はまだ酔っ払っていません。",
			2, 10,
			[][]int{
				{200, 5},
				{350, 3},
			},
			2},

		{"入力例3",
			3, 1000000,
			[][]int{
				{1000, 100},
				{1000, 100},
				{1000, 100},
			},
			-1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC189Bその1(tt.N, tt.X, tt.VP)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
