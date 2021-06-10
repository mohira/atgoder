package ABC_problemB

import (
	"strconv"
	"strings"
	"testing"
)

// [ABC137B - One Clue](https://atcoder.jp/contests/abc137/tasks/abc137_b)
func AnswerABC137Bその1(K, X int) string {
	var stoneRange []string

	for i := X - K + 1; i <= X+K-1; i++ {
		stoneRange = append(stoneRange, strconv.Itoa(i))
	}

	return strings.Join(stoneRange, " ")
}

func TestAnswerABC137Bその1(t *testing.T) {
	tests := []struct {
		name string
		K, X int
		want string
	}{
		{"入力例1", 3, 7, "5 6 7 8 9"},
		{"入力例2", 4, 0, "-3 -2 -1 0 1 2 3"},
		{"入力例3", 1, 100, "100"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC137Bその1(tt.K, tt.X)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
