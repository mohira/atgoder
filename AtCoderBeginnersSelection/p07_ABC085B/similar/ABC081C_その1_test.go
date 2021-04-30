package p06_ABC071B

import (
	"fmt"
	"os"
	"sort"
	"testing"
)

// [ABC081C - Not so Diverse](https://atcoder.jp/contests/abc081/tasks/arc086_a)
func AnswerABC081Cその1(N, K int, A []int) int {
	/*
		方針: 整数をK個選んで、それらのうち、どれかが書かれているボールの数を最大化する

		書き換えるべき個数(答え) == ボールの総数(N個) - 「書き換えなくてよいボールの数の最大」

		「書き換えなくてよいボールの数の最大」は「整数の出現頻度が多い方からK種類を選んで、それらK種類の整数の頻度の合計」を調べればOK

		ex: N=6, K=2, A=[1 1 2 5 5 5] の場合

		出現頻度順にソートすると

		[ 5 5 5 1 1 2]

		で、頻度が大きい方からK(=2)種類だけ選ぶと(=「5」と「1」を選ぶ)

		[ 5 5 5 1 1 2]
		 |o o o o o x| ← 5個

		よって、 6 - 5 = 1
	*/

	// 数字ごとの頻度情報を作成する
	var ballsFrequency = make(map[int]int)
	for _, a := range A {
		ballsFrequency[a]++
	}
	fmt.Fprintln(os.Stderr, ballsFrequency)

	// 数字の種類がK種類以下なら書き換え不要
	if len(ballsFrequency) <= K {
		return 0
	}

	// 頻度降順ソート
	// 「どの整数が」という情報は不要！ 出現頻度だけでいい
	frequencyDesc := make([]int, 0, len(ballsFrequency))
	for _, v := range ballsFrequency {
		frequencyDesc = append(frequencyDesc, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(frequencyDesc)))
	fmt.Fprintln(os.Stderr, frequencyDesc)

	// 頻度上位K種類のボールは書き換えなくてよい
	var ballsNotToRewrite int
	for _, f := range frequencyDesc[:K] {
		ballsNotToRewrite += f
	}

	return N - ballsNotToRewrite
}

func TestAnswerABC081Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		K    int
		A    []int
		want int
	}{
		{"おれおれケース: 整数順と頻度順が一致しないパターン", 6, 2, []int{1, 1, 2, 5, 5, 5}, 1},
		{"入力例1", 5, 2, []int{1, 1, 2, 2, 5}, 1},
		{"入力例2", 4, 4, []int{1, 1, 2, 2}, 0},
		{"入力例3", 10, 3, []int{5, 1, 3, 2, 4, 1, 1, 2, 3, 4}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC081Cその1(tt.N, tt.K, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
