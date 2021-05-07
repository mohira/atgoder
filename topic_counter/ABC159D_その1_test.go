package topic_counter

import (
	"fmt"
	"testing"
)

// [ABC159D - Banned K](https://atcoder.jp/contests/abc159/tasks/abc159_d)
func AnswerABC159Dその1(N int, A []int) string {
	bucket := make(map[int]int)
	for _, a := range A {
		bucket[a]++
	}

	// 全通り
	U := calcUniversalSet(bucket)

	var ans string

	for _, a := range A {
		// X個あったボールから1つを取り除くと、(X-1)通り だけ選択肢が減る
		ans += fmt.Sprintf("%d\n", U-(bucket[a]-1))
	}
	return ans
}

func calcUniversalSet(bucket map[int]int) int {
	count := 0
	for _, freq := range bucket {
		count += nC2(freq)
	}
	return count
}

func nC2(n int) int {
	return n * (n - 1) / 2
}

func TestAnswerABC159Dその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want string
	}{
		{"入力例1",
			5,
			[]int{1, 1, 2, 1, 2},
			"2\n2\n3\n2\n3\n",
		},

		{"入力例2",
			4,
			[]int{1, 2, 3, 4},
			"0\n0\n0\n0\n",
		},

		{"入力例3",
			5,
			[]int{3, 3, 3, 3, 3},
			"6\n6\n6\n6\n6\n",
		},
		{"入力例4",
			8,
			[]int{1, 2, 1, 4, 2, 1, 4, 1},
			"5\n7\n5\n7\n7\n5\n7\n5\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC159Dその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
