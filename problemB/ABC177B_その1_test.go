package problemB

import (
	"errors"
	"log"
	"testing"
)

// [ABC177B - Substring](https://atcoder.jp/contests/abc177/tasks/abc177_b)
func AnswerABC177Bその1(S, T string) int {
	var maxMatch int

	for i := 0; i < len(S)-len(T)+1; i++ {
		substr := S[i : i+len(T)]

		count, err := countMatchCharacters(substr, T)
		if err != nil {
			log.Fatal(err)
		}

		maxMatch = max(maxMatch, count)
	}

	return len(T) - maxMatch
}

// 2つ文字列が何文字一致しているかを数える
func countMatchCharacters(s1, s2 string) (int, error) {
	// 同じ文字数の前提
	if len(s1) != len(s2) {
		return 0, errors.New("s1 と s2 の文字数が違うよ")
	}

	var count int
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			count++
		}
	}

	return count, nil
}

func TestAnswerABC177Bその1(t *testing.T) {
	tests := []struct {
		name string
		S    string
		T    string
		want int
	}{
		{"入力例1",
			"cabacc",
			"abc",
			1,
		},
		{"入力例2",
			"codeforces",
			"atcoder",
			6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC177Bその1(tt.S, tt.T)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
