package topic_bit_brute_force

import (
	"strconv"
	"strings"
	"testing"
)

// [ABC045C - Bowls and Dishes](https://atcoder.jp/contests/abc045/tasks/abc045_c)
func AnswerABC045Cその1(S string) int {
	N := len(S) // Sの桁数

	var ans int

	// "+"を差し込めるスロットは、Sの桁数より1つ少ない
	for bit := 0; bit < (1 << (N - 1)); bit++ {
		var statement string
		for i := 0; i < N; i++ {
			s := string(S[i])
			statement += s
			if (bit>>i)&1 == 1 {
				// + を差し込む
				statement += "+"
			}
		}

		// "12+5" や "1+2+5" を計算する
		sum := 0
		for _, s := range strings.Split(statement, "+") {
			n, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			sum += n
		}
		ans += sum
	}

	return ans
}

func TestAnswerABC045Cその1(t *testing.T) {
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
			got := AnswerABC045Cその1(tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
