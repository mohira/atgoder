package topic_amari

import (
	"testing"
)

// [ABC055B - Training Camp](https://atcoder.jp/contests/abc055/tasks/abc055_b)
func AnswerABC055Bその1(N int) int {
	// 積abの余りは、a,bそれぞれの余りの積の余り
	// (ab mod X) = (a mod X) (b mod X) mod X
	// つまり、毎回、余りの計算をして良い
	const X = 1000000007
	var power = 1

	for i := 1; i <= N; i++ {
		power = (power * i) % X
	}

	return power
}

func TestAnswerABC055Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{"入力例1", 3, 6},
		{"入力例2", 10, 3628800},
		{"入力例3", 100000, 457992974},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC055Bその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
