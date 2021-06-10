package ABC_problemB

import (
	"testing"
)

// [ABC109B - Shiritori](https://atcoder.jp/contests/abc109/tasks/abc109_b)
func AnswerABC109Bその1(N int, W []string) string {
	bucket := make(map[string]int)

	// 複数回発言チェック
	for _, w := range W {
		if _, ok := bucket[w]; !ok {
			bucket[w]++
		} else {
			return "No"
		}
	}

	// 接続チェック
	for i := 0; i < N-1; i++ {
		now, next := W[i], W[i+1]

		tail := now[len(now)-1]
		head := next[0]
		if tail != head {
			return "No"
		}
	}

	return "Yes"
}

func TestAnswerABC109Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		W    []string
		want string
	}{
		{
			"入力例1",
			4,
			[]string{
				"hoge",
				"english",
				"hoge",
				"enigma",
			},
			"No",
		},
		{
			"入力例2",
			9,
			[]string{
				"basic",
				"c",
				"cpp",
				"php",
				"python",
				"nadesico",
				"ocaml",
				"lua",
				"assembly",
			},
			"Yes",
		},
		{
			"入力例3",
			3,
			[]string{
				"abc",
				"arc",
				"agc",
			},
			"No",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC109Bその1(tt.N, tt.W)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
