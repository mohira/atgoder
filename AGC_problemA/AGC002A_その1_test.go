package AGC_problemA

import (
	"testing"
)

// [AGC002A - Range Product](https://atcoder.jp/contests/agc002/tasks/agc002_a)
func AnswerAGC002その1(a, b int) string {
	if a <= 0 && 0 <= b {
		return "Zero"
	}

	if b < 0 {
		n := (-a) - (-b) + 1
		if n%2 == 0 {
			return "Positive"
		} else {
			return "Negative"
		}
	}

	return "Positive"
}

func TestAnswerAGC002その1(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want string
	}{
		{"入力例1", 1, 3, "Positive"},
		{"入力例2", -3, -1, "Negative"},
		{"入力例3", -1, 1, "Zero"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerAGC002その1(tt.a, tt.b)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
