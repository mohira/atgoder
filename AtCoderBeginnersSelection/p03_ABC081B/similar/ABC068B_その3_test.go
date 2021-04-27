package similar

import (
	"testing"
)

// [ABC068B - Break Number](https://atcoder.jp/contests/abc068/tasks/abc068_b)
func AnswerABC068Bその3(N int) int {
	// 全探索でシミュレーションする愚直な感じ
	var maxCount, count int
	var ans = 1
	for i := 1; i <= N; i++ {
		tmpNum := i
		count = 0

		for {
			if tmpNum%2 == 0 {
				count++
				tmpNum /= 2
			} else {
				break
			}
		}

		if maxCount < count {
			maxCount = count
			ans = i
		}
	}

	return ans
}

func TestAnswerABC068Bその3(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{"入力例1", 7, 4},
		{"入力例2", 32, 32},
		{"入力例3", 1, 1},
		{"入力例3", 99, 64},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC068Bその3(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}

		})

	}
}
