package problemC

import (
	"testing"
)

// [ABC176C - Step](https://atcoder.jp/contests/abc176/tasks/abc176_c)
func AnswerABC176Cその1(N int, A []int) int {
	var ans int

	for i := 0; i < N-1; i++ {
		now, next := A[i], A[i+1]

		// 次が小さい場合は、その場で補正していけばいい
		// その場で補正するのはもとの数列を破壊するのでイマイチかも
		if next < now {
			diff := now - next
			A[i+1] += diff
			ans += diff
		}
	}

	return ans
}

func TestAnswerABC176Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want int
	}{
		{"入力例1", 5, []int{2, 1, 5, 4, 3}, 4},
		{"入力例2", 5, []int{3, 3, 3, 3, 3}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC176Cその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
