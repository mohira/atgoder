package ABC_problemA

import (
	"testing"
)

// [ABC035A - Best Body](https://atcoder.jp/contests/abc035/tasks/abc035_a)
func AnswerABC035Aその1(W, H int) string {
	// 3W = 4H を満たすかどうかで判定したほうがかっこいいと思う
	if W%16 == 0 && H%9 == 0 {
		return "16:9"
	} else {
		return "4:3"
	}
}

func TestAnswerABC035Aその1(t *testing.T) {
	tests := []struct {
		name string
		W, H int
		want string
	}{
		{"入力例1", 4, 3, "4:3"},
		{"入力例2", 16, 9, "16:9"},
		{"入力例3", 28, 21, "4:3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC035Aその1(tt.W, tt.H)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
