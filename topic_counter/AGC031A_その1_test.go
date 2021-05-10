package topic_counter

import (
	"testing"
)

// [AGC031A - Colorful Subsequence](https://atcoder.jp/contests/agc031/tasks/agc031_a)
func AnswerAGC031Aその1(N int, S string) int {
	// 部分文字列として採用するときに場所は動かせないので、英文字ごとに独立して考えられる
	// すべての英文字について 「出現回数が1回以下」 を満たせば良い
	// ただし、 「部分文字列の長さが1以上」でないといけないことに注意
	//
	// ある英文字がn個あったとして、その英文字について次のパターンがある。
	//	 1. 部分文字列にその英文字を1つも選ばない ← 1通り == nC0
	//	 2. か 「その英文字を1つ選ぶ」がある ← n通り == nC1
	//
	// よって、英文字ごとのパターン数の積をとり、最後に1を引けばいい(すべて選ばないだと部分文字列にならないから)
	bucket := make(map[string]int)
	for _, c := range S {
		bucket[string(c)]++
	}

	pattern := 1
	for _, freq := range bucket {
		// nC0 + nC1
		pattern *= 1 + freq
	}

	// 何も選ばないケースを引き算する
	return pattern - 1
}

func TestAnswerAGC031Aその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		S    string
		want int
	}{
		{"入力例1", 2, "abcd", 15},
		{"入力例2", 3, "baa", 5},
		{"入力例3", 5, "abcab", 17},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerAGC031Aその1(tt.N, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
