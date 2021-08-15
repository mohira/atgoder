package ABC_problemC

import (
	"atgoder/lib"
	"fmt"
	"testing"
)

// [ABC064C - Colorful Leaderboard](https://atcoder.jp/contests/abc064/tasks/abc064_c)
func AnswerABC064Cその1(N int, A []int) string {
	bucket := make(map[string]int)
	wildCardCount := 0

	for _, a := range A {
		c := toColor(a)
		if c == "*" {
			wildCardCount++
		} else {
			bucket[c]++
		}
	}

	uniqColors := len(bucket)

	min := lib.Max(1, uniqColors)
	max := uniqColors + wildCardCount

	return fmt.Sprintf("%d %d", min, max)
}

func toColor(rate int) string {
	switch {
	case 1 <= rate && rate <= 399:
		return "灰"
	case 400 <= rate && rate <= 799:
		return "茶"
	case 800 <= rate && rate <= 1199:
		return "緑"
	case 1200 <= rate && rate <= 1599:
		return "水"
	case 1600 <= rate && rate <= 1999:
		return "青"
	case 2000 <= rate && rate <= 2399:
		return "黄"
	case 2400 <= rate && rate <= 2799:
		return "橙"
	case 2800 <= rate && rate <= 3199:
		return "赤"
	default:
		return "*"
	}
}

func TestAnswerABC064Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want string
	}{
		{"入力例1", 4, []int{2100, 2500, 2700, 2700}, "2 2"},
		{"入力例2", 5, []int{1100, 1900, 2800, 3200, 3200}, "3 5"},
		{"入力例3", 20, []int{800, 810, 820, 830, 840, 850, 860, 870, 880, 890, 900, 910, 920, 930, 940, 950, 960, 970, 980, 990}, "1 1"},

		{"自由に色を選べるレートしかない場合は最低1種類", 3, []int{3200, 3200, 3200}, "1 3"},
		{"色種類は超えてもOK", 9, []int{3200, 3200, 3200, 3200, 3200, 3200, 3200, 3200, 3200}, "1 9"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC064Cその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
