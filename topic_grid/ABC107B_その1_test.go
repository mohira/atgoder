package topic_grid

import (
	"testing"
)

// [ABC107B - Minesweeper](https://atcoder.jp/contests/abc107/tasks/abc107_b)
func AnswerABC107Bその1(H, W int, A []string) string {
	return ""
}

func TestAnswerABC107Bその1(t *testing.T) {
	tests := []struct {
		name string
		H, W int
		A    []string
		want string
	}{
		{"入力例1",
			4, 4,
			[]string{
				"##.#",
				"....",
				"##.#",
				".#.#",
			},
			"###\n###\n.##\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC107Bその1(tt.H, tt.W, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
