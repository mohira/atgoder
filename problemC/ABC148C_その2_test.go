package topic_amari

import (
	"testing"
)

// [ABC148C - Snack](https://atcoder.jp/contests/abc148/tasks/abc148_c)
func AnswerABC148Cその2(A, B int) int {
	// AxB は公倍数なので、B以下のAの倍数を探索して、Bの倍数かを調べればOK
	for i := 1; i <= B; i++ {
		if A*i%B == 0 {
			return A * i
		}
	}
	return 0
}

func TestAnswerABC148Cその2(t *testing.T) {
	tests := []struct {
		name string
		A, B int
		want int
	}{
		{"入力例1", 2, 3, 6},
		{"入力例2", 123, 456, 18696},
		{"入力例3", 100000, 99999, 9999900000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC148Cその2(tt.A, tt.B)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
