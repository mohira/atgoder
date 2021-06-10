package ABC_problemB

import (
	"math"
	"sort"
	"testing"
)

// [ABC026B - N重丸](https://atcoder.jp/contests/abc026/tasks/abc026_b)
func AnswerABC026Bその1(N int, R []int) float64 {
	sort.Sort(sort.Reverse(sort.IntSlice(R)))

	tmp := 0
	for i := 0; i < N; i++ {
		r := R[i]
		if i%2 == 0 {
			tmp += r * r
		} else {
			tmp -= r * r
		}
	}

	return float64(tmp) * math.Pi
}

func TestAnswerABC026Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		R    []int
		want float64
	}{
		{
			"入力例1",
			3,
			[]int{1, 2, 3},
			18.8495559215,
		},
		{
			"入力例2",
			6,
			[]int{15, 2, 3, 7, 6, 9},
			508.938009881546,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC026Bその1(tt.N, tt.R)

			if math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
