package similar

import (
	"sort"
	"testing"
)

// [ABC042B - 文字列大好きいろはちゃんイージー](https://atcoder.jp/contests/abc042/tasks/abc042_b)
func AnswerABC042Bその1(N, L int, S []string) string {
	sort.Strings(S)

	var ans string

	for _, s := range S {
		ans += s
	}

	return ans
}

func TestAnswerABC042Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		L    int
		S    []string
		want string
	}{
		{"入力例1", 3, 3, []string{"dxx", "axx", "cxx"}, "axxcxxdxx"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC042Bその1(tt.N, tt.L, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}

		})

	}
}
