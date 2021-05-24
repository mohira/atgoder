package problemB

import (
	"testing"
)

// [ABC125B - Resale](https://atcoder.jp/contests/abc125/tasks/abc125_b)
func AnswerABC125Bその2(N int, V, C []int) int {
	// N<=20 なので、ビット全探索でもいける
	var ans int
	for bit := 0; bit < (1 << N); bit++ {
		tmpValues := 0
		for i := 0; i < N; i++ {
			if (bit>>i)&1 == 1 {
				// フラグが立っているときの処理
				tmpValues += V[i] - C[i]
			}
		}

		ans = max(tmpValues, ans)
	}

	return ans
}

func TestAnswerABC125Bその2(t *testing.T) {
	tests := []struct {
		name string
		N    int
		V    []int
		C    []int
		want int
	}{
		{
			"入力例1",
			3,
			[]int{10, 2, 5},
			[]int{6, 3, 4},
			5,
		},
		{
			"入力例2",
			4,
			[]int{13, 21, 6, 19},
			[]int{11, 30, 6, 15},
			6,
		},
		{
			"入力例3",
			1,
			[]int{1},
			[]int{50},
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC125Bその2(tt.N, tt.V, tt.C)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
