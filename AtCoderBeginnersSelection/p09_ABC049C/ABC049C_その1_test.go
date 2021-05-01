package p08_ABC049C

import (
	"strings"
	"testing"
)

// [ABC049C - 白昼夢](https://atcoder.jp/contests/abc049/tasks/arc065_a)
func AnswerABC049Cその1(S string) string {
	var s1, s2, s3, s4 = "dream", "dreamer", "erase", "eraser"
	var divideCandidates = [4]string{reverse(s1), reverse(s2), reverse(s3), reverse(s4)}

	s := reverse(S)

	for {
		var matchWordAtLeastOne bool

		for _, candidate := range divideCandidates {
			if strings.HasPrefix(s, candidate) {
				s = strings.TrimPrefix(s, candidate)
				matchWordAtLeastOne = true
			}
		}

		// 候補のマッチが1つもない => 探索終了
		if !matchWordAtLeastOne {
			return "NO"
		}

		// sが空文字 == すべてうまくマッチした
		if s == "" {
			return "YES"
		}
	}
}

func reverse(s string) string {
	var newS string
	for i := len(s) - 1; i >= 0; i-- {
		newS += string(s[i])
	}
	return newS
}

func TestAnswerABC049Cその2(t *testing.T) {
	tests := []struct {
		name string
		S    string
		want string
	}{
		{"入力例1", "erasedream", "YES"},
		{"入力例2", "dreameraser", "YES"},
		{"入力例3", "dreamerer", "NO"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC049Cその1(tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
