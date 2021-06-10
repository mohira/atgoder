package ABC_problemB

import (
	"testing"
)

// [ABC103B - AcCepted](https://atcoder.jp/contests/abc103/tasks/abc103_b)
func AnswerABC103Bその1(S, T string) string {
	N := len(S)

	for j := 0; j < N; j++ {
		tmp := ""

		// ローテーション文字列の作成
		for i := 0; i < N; i++ {
			nextIdx := i + 1

			if nextIdx == N {
				tmp += string(S[0])
			} else {
				tmp += string(S[nextIdx])
			}
		}

		if tmp == T {
			return "Yes"
		}
		S = tmp
	}

	return "No"
}

func TestAnswerABC103Bその1(t *testing.T) {
	tests := []struct {
		name string
		S    string
		T    string
		want string
	}{
		{"入力例1", "kyoto", "tokyo", "Yes"},
		{"入力例2", "abc", "arc", "No"},
		{"入力例3", "aaaaaaaaaaaaaaab", "aaaaaaaaaaaaaaab", "Yes"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC103Bその1(tt.S, tt.T)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
