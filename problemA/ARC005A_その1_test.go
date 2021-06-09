package problemA

import (
	"strings"
	"testing"
)

// [ARC005A - 大好き高橋君](https://atcoder.jp/contests/arc005/tasks/arc005_1)
func AnswerARC005Aその1(N int, W []string) int {
	count := 0
	t1 := "TAKAHASHIKUN"
	t2 := "Takahashikun"
	t3 := "takahashikun"

	for i := 0; i < N; i++ {
		w := W[i]
		if i == N-1 {
			w = strings.TrimSuffix(W[i], ".")
		}

		if w == t1 || w == t2 || w == t3 {
			count++
		}
	}

	return count
}

func TestAnswerARC005Aその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		W    []string
		want int
	}{
		{"入力例1", 5, []string{"Takahashikun", "is", "not", "an", "eel."}, 1},
		{"入力例2", 5, []string{"TAKAHASHIKUN", "loves", "TAKAHASHIKUN", "and", "takahashikun."}, 3},
		{"入力例3", 6, []string{"He", "is", "not", "takahasikun", "but", "Takahashikun."}, 1},
		{"入力例4", 1, []string{"takahashikunTAKAHASHIKUNtakahashikun"}, 0},
		{"入力例5", 18, []string{"You", "should", "give", "Kabayaki", "to", "Takahashikun", "on", "July", "twenty", "seventh", "if", "you", "suspect", "that", "he", "is", "an", "eel."}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerARC005Aその1(tt.N, tt.W)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
