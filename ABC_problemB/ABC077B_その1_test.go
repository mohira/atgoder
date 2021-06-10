package ABC_problemB

import (
	"math"
	"testing"
)

// [ABC077B - ISU](https://atcoder.jp/contests/abc077/tasks/abc077_b)
func AnswerABC077Bその1(N int) int {
	// N = 1e+9 のとき √1e+9 = 31622.2 <= 31623
	// 999950884 = 31622**2 < 1e+9 < 31623**2 = 1000014129
	maxN := int(math.Sqrt(1e+9))

	var ans = 1
	for i := 1; i <= maxN; i++ {
		s := i * i
		if s <= N {
			ans = s
		} else {
			break
		}
	}
	return ans
}

func TestAnswerABC077Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{"入力例1", 10, 9},
		{"入力例2", 81, 81},
		{"入力例3", 271828182, 271821169},
		{"最大ケース", 1000000000, 999950884},
		{"最小ケース", 1, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC077Bその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
