package topic_brute_force

import (
	"testing"
)

// [ABC150B - 81](https://atcoder.jp/contests/abc150/tasks/abc150_b)
func AnswerABC150Bその1(N int, S string) int {
	// return strings.Count(S, "ABC") // これで一撃
	var count int

	for i := 0; i < N-2; i++ {
		if S[i:i+3] == "ABC" {
			count++
		}
	}
	return count
}

func TestAnswerABC150Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		S    string
		want int
	}{
		{"入力例1", 10, "ZABCDBABCQ", 2},
		{"入力例2", 19, "THREEONEFOURONEFIVE", 0},
		{"入力例3", 33, "ABCCABCBABCCABACBCBBABCBCBCBCABCB", 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC150Bその1(tt.N, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
