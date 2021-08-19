package ABC_problemB

import (
	"fmt"
	"testing"
)

// [ABC012B - 入浴時間](https://atcoder.jp/contests/abc012/tasks/abc012_2)
func AnswerABC012Bその1(N int) string {
	var hh, mm, ss int
	hh = N / 3600
	mm = (N / 60) % 60
	ss = N % 60

	return fmt.Sprintf("%02d:%02d:%02d", hh, mm, ss)
}

func TestAnswerABC012Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want string
	}{
		{"入力例1", 3661, "01:01:01"},
		{"入力例2", 86399, "23:59:59"},
		{"入力例3", 65, "00:01:05"},
		{"入力例4", 7265, "02:01:05"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC012Bその1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
