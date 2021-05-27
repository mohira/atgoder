package topic_grid

import (
	"fmt"
	"strings"
	"testing"
)

// [ABC062B - Picture Frame](https://atcoder.jp/contests/abc062/tasks/abc062_b)
func AnswerABC062Bその1(H, W int, S []string) string {
	var ans string
	ans += strings.Repeat("#", 1+W+1) + "\n"

	for _, s := range S {
		ans += fmt.Sprintf("#%s#\n", s)
	}

	ans += strings.Repeat("#", 1+W+1) + "\n"

	return ans
}

func TestAnswerABC062Bその1(t *testing.T) {
	tests := []struct {
		name string
		H    int
		W    int
		A    []string
		want string
	}{
		{"入力例1",
			2, 3,
			[]string{
				"abc",
				"arc",
			},
			"#####\n#abc#\n#arc#\n#####\n",
		},
		{"入力例2",
			1, 1,
			[]string{
				"z",
			},
			"###\n#z#\n###\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC062Bその1(tt.H, tt.W, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
