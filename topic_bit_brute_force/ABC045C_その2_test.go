package topic_bit_brute_force

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

// [ABC045C - Bowls and Dishes](https://atcoder.jp/contests/abc045/tasks/abc045_c)
func AnswerABC045Cその2(S string) int {
	var ans int

	n := len(S) - 1

	for bit := 0; bit < (1 << n); bit++ {
		expression := string(S[0])

		for i := 0; i < n; i++ {
			if (bit>>i)&1 == 1 {
				expression += "+"
			}
			expression += string(S[i+1])
		}

		fmt.Fprintln(os.Stderr, expression)

		for _, s := range strings.Split(expression, "+") {
			v, _ := strconv.Atoi(s)
			ans += v
		}
	}

	return ans
}

func TestAnswerABC045Cその2(t *testing.T) {
	tests := []struct {
		name string
		S    string
		want int
	}{
		{"入力例1", "125", 176},
		{"入力例2", "9999999999", 12656242944},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC045Cその2(tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
