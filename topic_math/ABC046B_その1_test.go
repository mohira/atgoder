package topic_math

import (
	"math"
	"testing"
)

// [ABC046B - AtCoDeerくんとボール色塗り](https://atcoder.jp/contests/abc046/tasks/abc046_b)
func AnswerABC046Bその1(N, K int) int {
	// 左から考える。1つ色を選んだら、次はその色だけが使えない
	//  ○   ○   ○
	// [r]      r
	//  g  [g]
	//  b   b   b
	n := float64(N)
	k := float64(K)

	f := k * math.Pow(k-1, n-1)
	return int(f)
}

func TestAnswerABC046Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		K    int
		want int
	}{
		{"入力例1", 2, 2, 2},
		{"入力例2", 1, 10, 10},
		{"入力例3", 10, 8, 322828856},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC046Bその1(tt.N, tt.K)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
