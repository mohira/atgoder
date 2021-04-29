package p06_ABC071B

import (
	"testing"
)

// [ABC071B - Not Found](https://atcoder.jp/contests/abc071/tasks/abc071_b)
func AnswerABC071Bその1(S string) string {
	// バケット法を素直に使う
	// runeのまま処理する
	const bucketSize = 'z' + 1
	var bucket [bucketSize]int
	for _, c := range S {
		bucket[c]++
	}

	for i := 'a'; i <= 'z'; i++ {
		if bucket[i] == 0 {
			return string(byte(i))
		}
	}

	return "None"
}

func TestAnswerABC071Bその1(t *testing.T) {
	tests := []struct {
		name string
		S    string
		want string
	}{
		{"入力例1", "atcoderregularcontest", "b"},
		{"入力例2", "abcdefghijklmnopqrstuvwxyz", "None"},
		{"入力例3", "fajsonlslfepbjtsaayxbymeskptcumtwrmkkinjxnnucagfrg", "d"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC071Bその1(tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
