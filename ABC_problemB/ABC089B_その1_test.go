package ABC_problemB

import (
	"testing"
)

// [ABC089B -  Hina Arare](https://atcoder.jp/contests/abc089/tasks/abc089_b)
func AnswerABC089Bその1(N int, S []string) string {
	bucket := make(map[string]int)
	for _, c := range S {
		bucket[c]++
	}

	if len(bucket) == 4 {
		return "Four"
	} else {
		return "Three"
	}
}

func TestAnswerABC089Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		S    []string
		want string
	}{
		{
			"入力例1",
			6,
			[]string{"G", "W", "Y", "P", "Y", "W"},
			"Four",
		},
		{
			"入力例2",
			9,
			[]string{"G", "W", "W", "G", "P", "W", "P", "G", "G"},
			"Three",
		},
		{
			"入力例3",
			8,
			[]string{"P", "Y", "W", "G", "Y", "W", "Y", "Y"},
			"Four",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC089Bその1(tt.N, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
