package topic_cumlative_sum

import (
	"testing"
)

// [ABC105D - Candy Distribution](https://atcoder.jp/contests/abc105/tasks/abc105_d)
func AnswerABC105Dその1(N, M int, A []int) int {
	cumsum := make([]int, N+1)
	for i := 0; i < N; i++ {
		cumsum[i+1] = cumsum[i] + A[i]
	}

	SmodM := make([]int, len(cumsum))
	for i := 0; i < len(cumsum); i++ {
		SmodM[i] = cumsum[i] % M
	}

	bucket := make(map[int]int)
	for _, v := range SmodM {
		bucket[v]++
	}

	var ans int
	for _, freq := range bucket {
		if 2 <= freq {
			ans += nC2(freq)
		}
	}

	return ans
}

func TestAnswerABC105Dその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		M    int
		AB   []int
		want int
	}{
		{"入力例1", 3, 2, []int{4, 1, 5}, 3},
		{"入力例2", 13, 17, []int{29, 7, 5, 7, 9, 51, 7, 13, 8, 55, 42, 9, 81}, 6},
		{"入力例3", 10, 400000000, []int{1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000}, 25},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC105Dその1(tt.N, tt.M, tt.AB)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
