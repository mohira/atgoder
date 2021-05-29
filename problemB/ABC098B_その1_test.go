package problemB

import (
	"atgoder/lib"
	"testing"
)

// [ABC098B - Cut and Count](https://atcoder.jp/contests/abc098/tasks/abc098_b)
func AnswerABC098Bその1(N int, S string) int {
	var maxCount = 0
	for i := 1; i < N; i++ {
		left, right := S[:i], S[i:]
		leftBucket, rightBucket := toRuneBucket(left), toRuneBucket(right)

		count := 0

		if len(leftBucket) < len(rightBucket) {
			for key := range rightBucket {
				if _, ok := leftBucket[key]; ok {
					count++
				}
			}
		} else {
			for key := range leftBucket {
				if _, ok := rightBucket[key]; ok {
					count++
				}
			}
		}

		maxCount = lib.Max(maxCount, count)
	}
	return maxCount
}

func toRuneBucket(s string) map[rune]int {
	bucket := make(map[rune]int)

	for _, c := range s {
		bucket[c]++
	}

	return bucket
}

func TestAnswerABC098Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		S    string
		want int
	}{
		{"入力例1", 6, "aabbca", 2},
		{"入力例2", 10, "aaaaaaaaaa", 1},
		{"入力例3", 45, "tgxgdqkyjzhyputjjtllptdfxocrylqfqjynmfbfucbir", 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC098Bその1(tt.N, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
