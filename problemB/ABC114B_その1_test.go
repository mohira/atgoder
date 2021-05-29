package problemB

import (
	"atgoder/lib"
	"math"
	"strconv"
	"testing"
)

// [ABC114B - Christmas Eve Eve](https://atcoder.jp/contests/abc114/tasks/abc114_b)
func AnswerABC114Bその1(S string) int {
	var minDiff = math.MaxInt64
	for i := 0; i <= len(S)-3; i++ {
		X, err := strconv.Atoi(S[i : i+3])

		if err != nil {
			panic(err)
		}

		absDiff := lib.AbsInt(X - 753)
		minDiff = lib.Min(minDiff, absDiff)
	}

	return minDiff
}

func TestAnswerABC114Bその1(t *testing.T) {
	tests := []struct {
		name string
		S    string
		want int
	}{
		{"入力例1", "1234567876", 34},
		{"入力例2", "35753", 0},
		{"入力例3", "1111111111", 642},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC114Bその1(tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
