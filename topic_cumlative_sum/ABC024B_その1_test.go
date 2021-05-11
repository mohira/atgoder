package topic_cumlative_sum

import (
	"testing"
)

// [ABC024B - 自動ドア](https://atcoder.jp/contests/abc024/tasks/abc024_b)
func AnswerABC024Bその1(N, T int, A []int) int {
	maxA := findMax(A)

	imos := make([]int, maxA+T+1)

	for _, a := range A {
		a--
		imos[a]++
		imos[a+T]-- // 時間なので、 +1 は不要
	}

	var cumsum = make([]int, len(imos)+1)
	for i := 0; i < len(imos); i++ {
		cumsum[i+1] = cumsum[i] + imos[i]
	}

	var ans int
	for _, s := range cumsum {
		if 0 < s {
			ans++
		}
	}
	return ans
}

func findMax(A []int) int {
	var max int
	for _, a := range A {
		if max < a {
			max = a
		}
	}
	return max
}

func TestAnswerABC024Bその1(t *testing.T) {
	tests := []struct {
		name string
		N, T int
		A    []int
		want int
	}{
		{"小さい例",
			3, 5,
			[]int{5, 15, 17},
			12},
		{"入力例1",
			5, 10,
			[]int{20, 100, 105, 217, 314},
			45},
		{"入力例2",
			10, 10,
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			19},
		{"入力例3",
			10, 100000,
			[]int{3, 31, 314, 3141, 31415, 314159, 400000, 410000, 500000, 777777},
			517253,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC024Bその1(tt.N, tt.T, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
