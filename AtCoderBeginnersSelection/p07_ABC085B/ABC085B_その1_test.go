package p06_ABC085B

import (
	"sort"
	"testing"
)

// [ABC085B - Kagami Mochi](https://atcoder.jp/contests/abc085/tasks/abc085_b)
func AnswerABC085Bその1(N int, D []int) int {
	// 重複を除いた数が答え
	sort.Sort(sort.Reverse(sort.IntSlice(D)))
	var ans int
	var now = -1 // 1<=dなのでおk

	for i := 0; i < N; i++ {
		d := D[i]
		if now != d {
			ans++
		}
		now = d
	}

	return ans
}

func TestAnswerABC085Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		D    []int
		want int
	}{
		{"入力例1", 4, []int{10, 8, 8, 6}, 3},
		{"入力例2", 3, []int{15, 15, 15}, 1},
		{"入力例3", 7, []int{50, 30, 50, 100, 50, 80, 30}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC085Bその1(tt.N, tt.D)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
