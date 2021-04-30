package p06_ABC071B

import (
	"math"
	"testing"
)

// [ABC091B - Counting Roads](https://atcoder.jp/contests/abc091/tasks/abc091_b)
func AnswerABC091Bその1(N int, S []string, M int, T []string) int {
	m := make(map[string]int)

	// sは加点対象
	for _, s := range S {
		m[s]++
	}
	// tは減点対象
	for _, t := range T {
		m[t]--
	}

	// 加点対象に存在しない言葉を言ったほうが良い場合もある
	return myMax(maxValueFromMap(m), 0)
}

func myMax(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func maxValueFromMap(m map[string]int) int {
	var maxValue = -math.MaxInt64
	for _, v := range m {
		if maxValue < v {
			maxValue = v
		}
	}

	return maxValue
}

func TestAnswerABC091Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		S    []string
		M    int
		T    []string
		want int
	}{
		{"入力例1", 3, []string{"apple", "orange", "apple"}, 1, []string{"grape"}, 2},
		{"入力例2", 3, []string{"apple", "orange", "apple"}, 5, []string{"apple", "apple", "apple", "apple", "apple"}, 1},
		{"入力例3 減点対象しかない場合", 1, []string{"voldemort"}, 10, []string{"voldemort", "voldemort", "voldemort", "voldemort", "voldemort", "voldemort", "voldemort", "voldemort", "voldemort", "voldemort"}, 0},
		{"入力例4", 6, []string{"red", "red", "blue", "yellow", "yellow", "red"}, 5, []string{"red", "red", "yellow", "green", "blue"}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC091Bその1(tt.N, tt.S, tt.M, tt.T)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
