package problemB

import (
	"fmt"
	"strings"
	"testing"
)

// [ABC031B - 運動管理](https://atcoder.jp/contests/abc031/tasks/abc031_b)
func AnswerABC031Bその1(L, H, N int, A []int) string {
	var s = make([]string, N)
	for i, a := range A {
		if L <= a && a <= H {
			s[i] = "0"
		}
		if a < L {
			s[i] = fmt.Sprintf("%d", L-a)
		}
		if H < a {
			s[i] = "-1"
		}
	}

	return strings.Join(s, "\n")
}

func TestAnswerABC031Bその1(t *testing.T) {
	tests := []struct {
		name string
		L, H int
		N    int
		A    []int
		want string
	}{
		{
			"入力例1",
			300, 400,
			3,
			[]int{240, 350, 480},
			"60\n0\n-1",
		},
		{
			"入力例2",
			50, 80,
			5,
			[]int{10000, 50, 81, 80, 0},
			"-1\n0\n-1\n0\n50",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC031Bその1(tt.L, tt.H, tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
