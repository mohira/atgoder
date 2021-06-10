package ABC_problemC_gray

import (
	"bytes"
	"testing"
)

// [ABC143C - Slimes](https://atcoder.jp/contests/abc143/tasks/abc143_c)
func AnswerABC143Cその1(N int, S string) int {
	var out bytes.Buffer

	for i := 0; i < N; i++ {
		for i != N-1 && S[i] == S[i+1] {
			i++
		}
		out.WriteByte(S[i])
	}

	return len(out.String())
}

func TestAnswerABC143Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		S    string
		want int
	}{
		{"入力例1", 10, "aabbbbaaca", 5},
		{"入力例2", 5, "aaaaa", 1},
		{"入力例3", 20, "xxzaffeeeeddfkkkkllq", 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC143Cその1(tt.N, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
