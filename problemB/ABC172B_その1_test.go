package problemB

import (
	"testing"
)

// [ABC172B - Minor Change](https://atcoder.jp/contests/abc172/tasks/abc172_b)
func AnswerABC172Bその1(S, T string) int {
	var count int

	for i := 0; i < len(S); i++ {
		if S[i] != T[i] {
			count++
		}
	}
	return count
}

func TestAnswerABC172Bその1(t *testing.T) {
	tests := []struct {
		name string
		S    string
		T    string
		want int
	}{
		{"入力例1",
			"cupofcoffee",
			"cupofhottea",
			4,
		},
		{"入力例2",
			"abcde",
			"bcdea",
			5,
		},
		{"入力例3",
			"apple",
			"apple",
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC172Bその1(tt.S, tt.T)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
