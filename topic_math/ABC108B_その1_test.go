package topic_math

import (
	"testing"
)

// [ABC108B - Ruined Square](https://atcoder.jp/contests/abc108/tasks/abc108_b)
func AnswerABC108Bその1(x1, y1, x2, y2 int) string {
	return ""
}

func TestAnswerABC108Bその1(t *testing.T) {
	tests := []struct {
		name           string
		x1, y1, x2, y2 int
		want           string
	}{
		{
			"入力例1",
			0, 0, 0, 1,
			"-1 1 - 0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC108Bその1(tt.x1, tt.y1, tt.x2, tt.y2)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
