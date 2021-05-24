package problemB

import (
	"testing"
)

// [ABC148B - Strings with the Same Length](https://atcoder.jp/contests/abc148/tasks/abc148_b)
func AnswerABC148Bその1(N int, S, T string) string {
	var ans string

	for i := 0; i < N; i++ {
		ans += string(S[i])
		ans += string(T[i])
	}

	return ans
}

func TestAnswerABC148Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		S, T string
		want string
	}{
		{"入力例1", 2, "ip", "cc", "icpc"},
		{"入力例2", 8, "hmhmnknk", "uuuuuuuu", "humuhumunukunuku"},
		{"入力例3", 5, "aaaaa", "aaaaa", "aaaaaaaaaa"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC148Bその1(tt.N, tt.S, tt.T)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
