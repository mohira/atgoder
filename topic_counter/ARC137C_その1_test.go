package topic_counter

import (
	"sort"
	"strings"
	"testing"
)

// [ABC137C - Green Bin](https://atcoder.jp/contests/abc137/tasks/abc137_c)
func AnswerABC137Cその1(N int, S []string) int {
	// 「文字列集合内に同じ文字列のペアがいくつ作れるか？」という問題に帰着させる
	// ex: [abc, abc, xyz, stu, stu, abc] → 5通り(abcで3通り + stuで2通り)
	//
	// アナグラムが成立するには
	// 	1. 文字の出現頻度が完全に一致 ← 「英文字と出現回数」をそのままKeyにはできないのでOUT
	//  2. ソートした文字列が一致 ← これならバケットのKeyにできる
	//
	// あとは 出現頻度2以上において、nC2を足していけばいい
	var count int

	bucket := make(map[string]int)

	for _, s := range S {
		sortedS := sortString(s)

		bucket[sortedS]++
	}

	for _, freq := range bucket {
		if 2 <= freq {
			count += freq * (freq - 1) / 2
		}
	}
	return count
}

func sortString(s string) string {
	stringSlice := strings.Split(s, "")
	sort.Strings(stringSlice)

	var sortedS string

	for _, c := range stringSlice {
		sortedS += c
	}

	return sortedS
}

func TestAnswerABC137Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		S    []string
		want int
	}{
		{"入力例1", 3, []string{"acornistnt", "peanutbomb", "constraint"}, 1},
		{"入力例2", 2, []string{"oneplustwo", "ninemodsix"}, 0},
		{"入力例3", 5, []string{"abaaaaaaaa", "oneplustwo", "aaaaaaaaba", "twoplusone", "aaaabaaaaa"}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC137Cその1(tt.N, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
