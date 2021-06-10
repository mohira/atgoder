package ABC_problemB

import (
	"fmt"
	"testing"
)

// [ABC173B - Judge Status Summary](https://atcoder.jp/contests/abc173/tasks/abc173_b)
func AnswerABC173Bその1(N int, S []string) string {
	bucket := make(map[string]int)
	for _, s := range S {
		bucket[s]++
	}

	var ans string
	for _, s := range []string{"AC", "WA", "TLE", "RE"} {
		ans += fmt.Sprintf("%s x %d\n", s, bucket[s])
	}

	return ans
}

func TestAnswerABC173Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		S    []string
		want string
	}{
		{"入力例1",
			4, []string{"AC", "TLE", "AC", "AC", "WA", "TLE"},
			"AC x 3\nWA x 1\nTLE x 2\nRE x 0\n",
		},
		{"入力例2",
			10, []string{"AC", "AC", "AC", "AC", "AC", "AC", "AC", "AC", "AC", "AC"},
			"AC x 10\nWA x 0\nTLE x 0\nRE x 0\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC173Bその1(tt.N, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
