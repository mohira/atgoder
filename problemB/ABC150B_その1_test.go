package problemB

import (
	"testing"
)

// [ABC150B - Count ABC](https://atcoder.jp/contests/abc150/tasks/abc150_b)
func AnswerABC150Bその1(N int, S string) int {
	var count int
	for i := 0; i <= N-3; i++ {
		if S[i:i+3] == "ABC" {
			count++
		}
	}

	return count
	//return strings.Count(S, "ABC")
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
