package p06_ABC071B

import (
	"fmt"
	"testing"
)

// [ABC061B - Counting Roads](https://atcoder.jp/contests/abc061/tasks/abc061_b)
func AnswerABC061Bその1(N, M int, AB [][2]int) string {
	// すべての数の登場回数を数えればOK
	const nMax = 50
	var bucket [nMax + 1]int

	for i := 0; i < M; i++ {
		a, b := AB[i][0], AB[i][1]
		bucket[a]++
		bucket[b]++
	}

	var ans string
	for i := 1; i <= N; i++ {
		ans += fmt.Sprintf("%d\n", bucket[i])
	}

	return ans
}

func TestAnswerABC061Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		M    int
		AB   [][2]int
		want string
	}{
		{"入力例1",
			4,
			3,
			[][2]int{
				{1, 2},
				{2, 3},
				{1, 4},
			},
			"2\n2\n1\n1\n",
		},
		{"入力例2",
			2,
			5,
			[][2]int{
				{1, 2},
				{2, 1},
				{1, 2},
				{2, 1},
				{1, 2},
			},
			"5\n5\n",
		},
		{"入力例3",
			8,
			8,
			[][2]int{
				{1, 2},
				{3, 4},
				{1, 5},
				{2, 8},
				{3, 7},
				{5, 2},
				{4, 1},
				{6, 8},
			},
			"3\n3\n2\n2\n2\n1\n1\n2\n",
		},
		{"どこからも接続されていない都市がある場合",
			3,
			1,
			[][2]int{
				{1, 2},
			},
			"1\n1\n0\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC061Bその1(tt.N, tt.M, tt.AB)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
