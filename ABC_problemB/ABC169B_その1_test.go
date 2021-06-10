package ABC_problemB

import (
	"testing"
)

// [ABC169B - Multiplication 2](https://atcoder.jp/contests/abc169/tasks/abc169_b)
func AnswerABC169Bその1(N int, A []int64) int64 {
	// 0が含まれていたら答えは0
	for _, a := range A {
		if a == 0 {
			return 0
		}
	}

	var product int64 = 1

	for _, a := range A {
		// 1e+18 < a * product と同値
		if 1e+18/product < a {
			return -1
		} else {
			product *= a
		}
	}

	return product
}

func TestAnswerABC169Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int64
		want int64
	}{
		{"入力例1: 積そのままでOK", 2, []int64{1000000000, 1000000000}, 1000000000000000000},
		{"入力例2: 積が10^18を超えてしまうが、オーバーフローはしない", 3, []int64{101, 9901, 999999000001}, -1},
		{"入力例3: 途中で10^18を超えるが、最後に0があるので答えは0になるパターン", 31, []int64{4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9, 3, 2, 3, 8, 4, 6, 2, 6, 4, 3, 3, 8, 3, 2, 7, 9, 5, 0}, 0},
		{"オーバーフローするケース", 2, []int64{4294967296, 4294967296}, -1},
		{"オーバーフローするケース", 2, []int64{4294967296, 4294967296}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC169Bその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
