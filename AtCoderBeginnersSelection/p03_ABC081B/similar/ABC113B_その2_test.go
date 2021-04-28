package similar

import (
	"testing"
)

// [ABC113B - Palace](https://atcoder.jp/contests/abc113/tasks/abc113_b)
func AnswerABC113Bその2(T int, A int, H []int) int {
	// 気温を1000倍すれば、intの世界だけでやっていける
	var ansIdx int
	var minDiff = 1_000_000_000

	for i, h := range H {
		avgTemperature := T*1000 - h*6

		diff := 1000*A - avgTemperature
		if diff < 0 {
			// math.Absを使わないので自力
			diff = -diff
		}

		if diff < minDiff {
			ansIdx = i + 1
			minDiff = diff
		}
	}

	return ansIdx
}

func TestAnswerABC113Bその2(t *testing.T) {
	tests := []struct {
		name string
		N    int
		T    int
		A    int
		H    []int
		want int
	}{
		{"入力例1", 2, 12, 5, []int{1000, 2000}, 1},
		{"入力例2", 3, 21, -11, []int{81234, 94124, 52141}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC113Bその2(tt.T, tt.A, tt.H)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}

		})

	}
}
