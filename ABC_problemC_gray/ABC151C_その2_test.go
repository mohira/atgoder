package ABC_problemC_gray

import (
	"fmt"
	"testing"
)

// [ABC151C - Welcome to AtCoder](https://atcoder.jp/contests/abc151/tasks/abc151_c)
func AnswerABC151Cその2(N, M int, P []int, S []string) string {
	// 余計なバケットを挟まず実装する感じ
	acFlags := make([]bool, N)
	waCounts := make([]int, N)

	for i := 0; i < M; i++ {
		problemId := P[i] - 1
		result := S[i]
		if acFlags[problemId] {
			continue
		}

		if result == "AC" {
			acFlags[problemId] = true
		} else {
			waCounts[problemId]++
		}
	}

	var countAC, countPenalty int
	for i := 0; i < N; i++ {
		if acFlags[i] {
			countAC++
			countPenalty += waCounts[i]
		}
	}

	return fmt.Sprintf("%d %d", countAC, countPenalty)
}

func TestAnswerABC151Cその2(t *testing.T) {
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
			got := AnswerABC151Cその2(tt.N, tt.M, tt.P, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
