package ABC_problemB

import (
	"testing"
)

// [ABC165B - 1%](https://atcoder.jp/contests/abc165/tasks/abc165_b)
func AnswerABC165Bその1(X int) int {
	depositAmount := 100

	for i := 1; i <= 3760; i++ {
		// 1.01 を掛けるは誤差を生む
		// 利子1%で切り捨てなので、そもそも 預金額/100 を足すほうがいい
		depositAmount += depositAmount / 100
		if X <= depositAmount {
			return i
		}
	}

	return -1
}

func TestAnswerABC165Bその1(t *testing.T) {
	tests := []struct {
		name string
		X    int
		want int
	}{
		{"入力例1", 103, 3},
		{"入力例2", 1000000000000000000, 3760},
		{"入力例3", 1333333333, 1706},
		{"after_contest", 974755271730884810, 3758},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC165Bその1(tt.X)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
