package problemB

import (
	"testing"
)

// [ABC044B - 美しい文字列](https://atcoder.jp/contests/abc044/tasks/abc044_b)
func AnswerABC044Bその1(w string) string {
	bucket := make(map[rune]int)

	for _, c := range w {
		bucket[c]++
	}

	var can = true
	for _, freq := range bucket {
		if freq%2 != 0 {
			can = false
		}
	}

	if can {
		return "Yes"
	} else {
		return "No"
	}
}

func TestAnswerABC044Bその1(t *testing.T) {
	tests := []struct {
		name string
		w    string
		want string
	}{
		{"入力例1", "abaccaba", "Yes"},
		{"入力例2", "hthth", "No"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC044Bその1(tt.w)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
