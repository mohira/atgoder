package problemC

import (
	"fmt"
	"testing"
)

// [ABC151C - Welcome to AtCoder](https://atcoder.jp/contests/abc151/tasks/abc151_c)
func AnswerABC151Cその1(N, M int, P []int, S []string) string {
	// 考えやすくするために、問題ごとのresultスライスを用意する感じ
	bucket := make(map[int][]string)
	for i := 0; i < M; i++ {
		problemId := P[i] - 1
		result := S[i]
		bucket[problemId] = append(bucket[problemId], result)
	}

	countAC := 0
	countPenalty := 0

	for _, results := range bucket {
		isAC := false // 最終的にACかどうか

		countWA := 0
		for _, result := range results {
			if !isAC {
				if result == "AC" {
					isAC = true
				} else {
					countWA++
				}
			}
		}

		// 最終的にACだったときだけ正答数とペナルティ数は変動する
		if isAC {
			countAC++
			countPenalty += countWA
		}
	}

	return fmt.Sprintf("%d %d", countAC, countPenalty)
}

func TestAnswerABC151Cその1(t *testing.T) {
	tests := []struct {
		name string
		N, M int
		P    []int
		S    []string
		want string
	}{
		{
			"入力例1",
			2, 5,
			[]int{1, 1, 2, 2, 2},
			[]string{"WA", "AC", "WA", "AC", "WA"},
			"2 2",
		},
		{
			"入力例2",
			100000, 3,
			[]int{7777, 7777, 7777},
			[]string{"AC", "AC", "AC"},
			"1 0",
		},
		{
			"入力例3",
			6, 0,
			[]int{},
			[]string{},
			"0 0",
		},

		{
			"ACしていないので、いくらWAしててもペナルティにはカウントされない",
			1, 3,
			[]int{1, 1, 1},
			[]string{"WA", "WA", "WA"},
			"0 0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC151Cその1(tt.N, tt.M, tt.P, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
