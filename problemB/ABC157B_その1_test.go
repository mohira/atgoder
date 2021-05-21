package problemB

import (
	"testing"
)

// [ABC157B - Bingo](https://atcoder.jp/contests/abc157/tasks/abc157_b)
func AnswerABC157Bその1(A [][]int, N int, B []int) string {
	var flag [3][3]bool

	for _, b := range B {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if A[i][j] == b {
					flag[i][j] = true
				}
			}
		}
	}

	yoko1 := flag[0][0] && flag[0][1] && flag[0][2]
	yoko2 := flag[1][0] && flag[1][1] && flag[1][2]
	yoko3 := flag[2][0] && flag[2][1] && flag[2][2]

	tate1 := flag[0][0] && flag[1][0] && flag[2][0]
	tate2 := flag[0][1] && flag[1][1] && flag[2][1]
	tate3 := flag[0][2] && flag[1][2] && flag[2][2]

	naname1 := flag[0][0] && flag[1][1] && flag[2][2]
	naname2 := flag[0][2] && flag[1][1] && flag[2][0]

	if yoko1 || yoko2 || yoko3 || tate1 || tate2 || tate3 || naname1 || naname2 {
		return "Yes"
	}

	return "No"
}

func TestAnswerABC157Bその1(t *testing.T) {
	tests := []struct {
		name string
		A    [][]int
		N    int
		B    []int
		want string
	}{
		{
			"入力例1",
			[][]int{
				{84, 97, 66},
				{79, 89, 11},
				{61, 59, 7},
			},
			7,
			[]int{89, 7, 87, 79, 24, 84, 30},
			"Yes",
		},
		{
			"入力例2",
			[][]int{
				{41, 7, 46},
				{26, 89, 2},
				{78, 92, 8},
			},
			5,
			[]int{6, 45, 16, 57, 17},
			"No",
		},
		{
			"入力例3",
			[][]int{
				{60, 88, 34},
				{92, 41, 43},
				{65, 73, 48},
			},
			10,
			[]int{60, 43, 88, 11, 48, 73, 65, 41, 92, 34},
			"Yes",
		},
		{
			"",
			[][]int{
				{72, 28, 21},
				{35, 74, 30},
				{2, 81, 46},
			},
			7,
			[]int{85, 30, 51, 21, 46, 47},
			"Yes",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC157Bその1(tt.A, tt.N, tt.B)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
