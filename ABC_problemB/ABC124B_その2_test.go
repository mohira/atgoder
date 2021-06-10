package ABC_problemB

import (
	"testing"
)

// [ABC124B - Great Ocean View](https://atcoder.jp/contests/abc124/tasks/abc124_b)
func AnswerABC124Bその2(N int, H []int) int {
	var count = 1
	for i := 1; i < N; i++ {
		can := true
		for j := 0; j < i; j++ {

			if H[i] < H[j] {
				can = false
			}
		}

		if can {
			count++
		}
	}
	return count
}

func TestAnswerABC124Bその2(t *testing.T) {
	tests := []struct {
		name string
		N    int
		H    []int
		want int
	}{
		{"入力例1", 4, []int{6, 5, 6, 8}, 3},
		{"入力例2", 5, []int{4, 5, 3, 5, 4}, 3},
		{"入力例3", 5, []int{9, 5, 6, 8, 4}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC124Bその2(tt.N, tt.H)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
