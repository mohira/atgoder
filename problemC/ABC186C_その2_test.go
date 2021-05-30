package topic_amari

import (
	"fmt"
	"strings"
	"testing"
)

// [ABC186C - Unlucky 7](https://atcoder.jp/contests/abc186/tasks/abc186_c)
func AnswerABC186Cその2(N int) int {
	// 基数変換を文字列で扱う
	var count int
	for i := 1; i <= N; i++ {
		s8 := fmt.Sprintf("%o", i)
		s10 := fmt.Sprintf("%d", i)

		if strings.Contains(s8, "7") || strings.Contains(s10, "7") {
			continue
		}

		count++
	}

	return count
}

func TestAnswerABC186Cその2(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{"入力例1", 20, 17},
		{"入力例2", 100000, 30555},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC186Cその2(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
