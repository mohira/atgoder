package problemB

import (
	"testing"
)

// [ABC059B - Comparison](https://atcoder.jp/contests/abc059/tasks/abc059_b)
func AnswerABC059Bその1(A, B string) string {
	// 先頭が0でないので、桁数が不一致なら、大or小はわかる
	nA := len(A)
	nB := len(B)

	if nA < nB {
		return "LESS"
	}

	if nA > nB {
		return "GREATER"
	}

	// 桁数が一致するなら文字列比較で決まる
	if A < B {
		return "LESS"
	}
	if A > B {
		return "GREATER"
	}

	return "EQUAL"
}

func TestAnswerABC059Bその1(t *testing.T) {
	tests := []struct {
		name string
		A, B string
		want string
	}{
		{"入力例1", "36", "24", "GREATER"},
		{"入力例2", "850", "3777", "LESS"},
		{"入力例3", "9720246", "22516266", "LESS"},
		{"入力例4", "123456789012345678901234567890", "234567890123456789012345678901", "LESS"},
		{"等しいケース", "1", "1", "EQUAL"},
		{"大きい", "360", "24", "GREATER"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC059Bその1(tt.A, tt.B)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
