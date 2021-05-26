package topic_math

import (
	"math"
	"testing"
)

// [ABC097B - Exponential](https://atcoder.jp/contests/abc097/tasks/abc097_b)
func AnswerABC097Bその1(X int) int {
	// べき乗数のバケット
	bucket := make(map[int]int)
	bucket[1] = 1 // 1のべき乗はあらかじめカウント

	// b^p を探索
	// 1000以下になる「べき乗数」の中で
	//  bがもっとも小さくなるのは、b=2; 31^2(=961) これは入力例3からわかる。雑に考えればX以下でもいいと思う。
	//  pが最も大きくなるのは、p=9; (2^9=512, 2^10=1024)
	for b := 2; b <= 31; b++ {
		for p := 2; p <= 9; p++ {
			tmp := int(math.Pow(float64(b), float64(p)))

			if tmp <= X {
				bucket[tmp] = 1
			}
		}
	}

	var ans int
	for value := range bucket {
		ans = max(ans, value)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func TestAnswerABC097Bその1(t *testing.T) {
	tests := []struct {
		name string
		X    int
		want int
	}{
		{"入力例1", 10, 9},
		{"入力例2", 1, 1},
		{"入力例3", 999, 961},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC097Bその1(tt.X)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
