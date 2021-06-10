package ABC_problemB

import (
	"testing"
)

// [ABC145B - Echo](https://atcoder.jp/contests/abc145/tasks/abc145_b)
func AnswerABC145Bその1(N int, S string) string {
	// 文字の長さは偶数でないといけない
	if N%2 != 0 {
		return "No"
	}

	left, right := S[:len(S)/2], S[len(S)/2:]

	if left == right {
		return "Yes"
	} else {
		return "No"
	}
}

func TestAnswerABC145Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		S    string
		want string
	}{
		{"入力例1", 6, "abcabc", "Yes"},
		{"入力例2", 6, "abcadc", "No"},
		{"入力例3", 1, "z", "No"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC145Bその1(tt.N, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
