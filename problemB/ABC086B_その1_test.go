package problemB

import (
	"fmt"
	"strconv"
	"testing"
)

// [ABC086B - 1 21](https://atcoder.jp/contests/abc086/tasks/abc086_b)
func AnswerABC086Bその1(a, b int) string {
	n, err := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	if err != nil {
		panic(err)
	}

	// 最大で、 100100^0.5 = 316.3 < 317
	for i := 1; i <= 317; i++ {
		if i*i == n {
			return "Yes"
		}
	}

	return "No"
}

func TestAnswerABC086Bその1(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want string
	}{
		{"入力例1", 1, 21, "Yes"},
		{"入力例2", 100, 100, "No"},
		{"入力例3", 12, 10, "No"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC086Bその1(tt.a, tt.b)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
