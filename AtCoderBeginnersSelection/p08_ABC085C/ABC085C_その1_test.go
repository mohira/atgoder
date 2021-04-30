package p08_ABC085C

import (
	"fmt"
	"testing"
)

// [ABC085C - Otoshidama](https://atcoder.jp/contests/abc085/tasks/arc086_a)
func AnswerABC085Cその1(N, Y int) string {
	for i := 0; i <= N; i++ {
		for j := 0; j <= N; j++ {
			k := N - i - j

			amount := 10000*i + 5000*j + 1000*k
			if k >= 0 && amount == Y {
				return fmt.Sprintf("%d %d %d", i, j, k)
			}
		}
	}
	return "-1 -1 -1"
}

func TestAnswerABC085Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		Y    int
		want string
	}{
		{"入力例1", 9, 45000, "0 9 0"},
		{"入力例2", 20, 196000, "-1 -1 -1"},
		{"入力例3", 1000, 1234000, "2 54 944"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC085Cその1(tt.N, tt.Y)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
