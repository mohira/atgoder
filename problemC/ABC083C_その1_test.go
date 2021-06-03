package problemC

import (
	"testing"
)

// [ABC083C - Multiple Gift](https://atcoder.jp/contests/abc083/tasks/abc083_c)
func AnswerABC083Cその1(X, Y int) int {
	// 「A[i+1] が A[i]の倍数」なので、2倍していくのがベスト
	// 2^60 = 1.1529215e+18 > 1e+18 なので、計算量は問題ない
	var count = 1 // X を初期カウント

	for X*2 <= Y {
		X *= 2
		count++
	}

	return count
}

func TestAnswerABC083Cその1(t *testing.T) {
	tests := []struct {
		name string
		X, Y int
		want int
	}{
		{"入力例1", 3, 20, 3},
		{"入力例2", 25, 100, 3},
		{"入力例3", 314159265, 358979323846264338, 31},
		{"最大ケース", 1, int(1e+18), 60},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC083Cその1(tt.X, tt.Y)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
