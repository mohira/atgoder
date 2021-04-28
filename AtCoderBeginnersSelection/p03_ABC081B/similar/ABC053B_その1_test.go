package similar

import (
	"strings"
	"testing"
)

// [ABC053B - A to Z String](https://atcoder.jp/contests/abc053/tasks/abc053_b)
func AnswerABC053Bその1(s string) int {
	// strings.Indexで登場するindexをしらべる
	// 最初のA → そのまま
	// 最後のZ → 反転した文字列の先頭のZ
	var sReverse string

	for i := len(s) - 1; i >= 0; i-- {
		sReverse += string(s[i])
	}
	lastIndexZ := (len(s) - 1) - strings.Index(sReverse, "Z")

	firstIndexA := strings.Index(s, "A")

	return lastIndexZ - firstIndexA + 1
}

func TestAnswerABC053Bその1(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"入力例1", "QWERTYASDFZXCV", 5},
		{"入力例2: ZがAより先に出現する場合", "ZABCZ", 4},
		{"入力例3: 末尾Zが連続している場合", "HASFJGHOGAKZZFEGA", 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC053Bその1(tt.s)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}

		})

	}
}
