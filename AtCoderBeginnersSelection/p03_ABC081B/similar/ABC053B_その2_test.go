package similar

import (
	"strings"
	"testing"
)

// [ABC053B - A to Z String](https://atcoder.jp/contests/abc053/tasks/abc053_b)
func AnswerABC053Bその2(s string) int {
	return findIndexLastZ(s) - findIndexFirstA(s) + 1
}

func findIndexFirstA(s string) int {
	return strings.Index(s, "A")
}

func findIndexLastZ(s string) int {
	var index = -1

	for i := 0; i < len(s); i++ {
		if string(s[i]) == "Z" {
			index = i
		}
	}

	return index
}

func TestAnswerABC053Bその2(t *testing.T) {
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
			got := AnswerABC053Bその2(tt.s)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
