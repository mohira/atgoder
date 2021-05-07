package similar

import (
	"sort"
	"testing"
)

// [ABC082B - Two Anagrams](https://atcoder.jp/contests/abc082/tasks/abc082_b)
// 文字列の比較演算で解決する
func AnswerABC082Bその2(s string, t string) string {
	if sortStringAsc(s) < sortStringDesc(t) {
		return "Yes"
	} else {
		return "No"
	}
}

func toIntSliceFromString(s string) []int {
	var slice []int

	for _, c := range s {
		slice = append(slice, int(c))
	}
	return slice
}

func toStringFromIntSlice(tmp []int) string {
	var s string

	for _, c := range tmp {
		s += string(uint8(c))
	}

	return s
}

func sortStringAsc(s string) string {
	intSlice := toIntSliceFromString(s)

	sort.Ints(intSlice)

	return toStringFromIntSlice(intSlice)

}

func sortStringDesc(t string) string {
	intSlice := toIntSliceFromString(t)

	sort.Sort(sort.Reverse(sort.IntSlice(intSlice)))

	return toStringFromIntSlice(intSlice)
}

func TestAnswerABC082Bその2(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		// N < M の場合
		{"入力例1", "yx", "axy", "Yes"},
		{"入力例4", "w", "ww", "Yes"},
		{"入力例3", "cd", "abc", "No"},

		// N = M の場合
		{"入力例5: 同じ文字列はNo", "zzz", "zzz", "No"},

		// N > M の場合
		{"入力例2", "ratcode", "atlas", "Yes"},
		{"N>Mかつ同一文字の連続はNo", "www", "ww", "No"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC082Bその2(tt.s, tt.t)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
