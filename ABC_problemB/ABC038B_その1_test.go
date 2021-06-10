package ABC_problemB

import (
	"testing"
)

// [ABC038B - ディスプレイ](https://atcoder.jp/contests/abc038/tasks/abc038_b)
func AnswerABC038Bその1(H1, W1, H2, W2 int) string {
	if H1 == H2 || H1 == W2 || W1 == H2 || W1 == W2 {
		return "YES"
	} else {
		return "NO"
	}
}

func TestAnswerABC038Bその1(t *testing.T) {
	tests := []struct {
		name           string
		H1, W1, H2, W2 int
		want           string
	}{
		{"入力例1", 1080, 1920, 1080, 1920, "YES"},
		{"入力例2", 1080, 1920, 1920, 1080, "YES"},
		{"入力例3", 334, 668, 668, 1002, "YES"},
		{"入力例4", 100, 200, 300, 150, "NO"},
		{"入力例5", 120, 120, 240, 240, "NO"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC038Bその1(tt.H1, tt.W1, tt.H2, tt.W2)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
