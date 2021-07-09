package ARC_problemA

import (
	"math"
	"testing"
)

// [ARC015A - Celsius と Fahrenheit](https://atcoder.jp/contests/arc015/tasks/arc015_1)
func AnswerARC015Aその1(n int) float64 {
	return 32 + 9*float64(n)/5
}

func TestAnswerARC015Aその1(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want float64
	}{
		{"入力例1", 10, 50},
		{"入力例2", 33, 91.4},
		{"入力例3", -100, -148},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerARC015Aその1(tt.n)

			if math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
