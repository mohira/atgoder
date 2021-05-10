package topic_cumlative_sum

import (
	"testing"
)

// [AGC023 A Zero-Sum Ranges](https://atcoder.jp/contests/agc023/tasks/agc023_a)
func AnswerAGC023Aその1(N int, A []int) int {
	// 累積和が等しい値を2つ選ぶと総和が0になる
	cumsum := []int{0}
	for i := 0; i < N; i++ {
		cumsum = append(cumsum, cumsum[i]+A[i])
	}

	bucket := make(map[int]int)
	for _, s := range cumsum {
		bucket[s]++
	}

	var ans int

	for _, freq := range bucket {
		if 2 <= freq {
			ans += nC2(freq)
		}
	}

	return ans
}

func nC2(n int) int {
	return n * (n - 1) / 2
}

func TestAnswerAGC023Aその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want int
	}{
		{"入力例1", 6, []int{1, 3, -4, 2, 2, -2}, 3},
		{"入力例2", 7, []int{1, -1, 1, -1, 1, -1, 1}, 12},
		{"入力例3", 5, []int{1, -2, 3, -4, 5}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerAGC023Aその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
