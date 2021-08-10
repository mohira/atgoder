package ABC_problemB

import (
	"testing"
)

// [ABC008B - 投票](https://atcoder.jp/contests/abc008/tasks/abc008_2)
func AnswerABC008Bその1(N int, S []string) string {
	bucket := make(map[string]int)

	for _, s := range S {
		bucket[s]++
	}

	var ans string
	var maxVote int

	for name, vote := range bucket {
		if maxVote < vote {
			maxVote = vote
			ans = name
		}
	}

	return ans
}

func TestAnswerABC008Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		S    []string
		want string
	}{
		{"入力例1",
			4,
			[]string{"taro", "jiro", "taro", "saburo"},
			"taro",
		},
		{"入力例2",
			1,
			[]string{"takahashikun"},
			"takahashikun",
		},
		{"入力例3",
			9,
			[]string{"a", "b", "c", "c", "b", "c", "b", "d", "e"},
			"b",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC008Bその1(tt.N, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
