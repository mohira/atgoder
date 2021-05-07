package topic_counter

import (
	"strconv"
	"testing"
)

// [ABC171D - Replacing](https://atcoder.jp/contests/abc171/tasks/abc171_d)
func AnswerABC171Dその1(N int, A []int, Q int, BC [][2]int) string {
	// 並び順に無関係 → バケット化
	// 差分「更新」 → 「変更する」ではなく、なにかの頻度を減らして、何かの頻度を増やす ←
	bucket := make(map[int]int)

	for _, a := range A {
		bucket[a]++
	}

	sum := sumBucket(bucket)

	var ans string
	for _, bc := range BC {
		b, c := bc[0], bc[1]

		sum += (c - b) * bucket[b]

		bucket[c] += bucket[b]
		bucket[b] = 0

		ans += strconv.Itoa(sum) + "\n"
	}

	return ans
}

func sumBucket(bucket map[int]int) int {
	var s int
	for num, freq := range bucket {
		s += num * freq
	}

	return s
}

func TestAnswerABC171Dその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		Q    int
		BC   [][2]int
		want string
	}{
		{"入力例1",
			4,
			[]int{1, 2, 3, 4},
			3,
			[][2]int{
				{1, 2},
				{3, 4},
				{2, 4},
			},
			"11\n12\n16\n",
		},
		{"入力例2",
			4,
			[]int{1, 1, 1, 1},
			3,
			[][2]int{
				{1, 2},
				{2, 1},
				{3, 5},
			},
			"8\n4\n4\n",
		},
		{"入力例3",
			2,
			[]int{1, 2},
			3,
			[][2]int{
				{1, 100},
				{2, 100},
				{100, 1000},
			},
			"102\n200\n2000\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC171Dその1(tt.N, tt.A, tt.Q, tt.BC)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
