package ABC_problemC_brown

import (
	"math"
	"strconv"
	"testing"
)

// [ABC157C - Guess The Number](https://atcoder.jp/contests/abc157/tasks/abc157_c)
func AnswerABC157Cその2(N, M int, SC [][]int) int {
	// 制約がゆるいので、全探索して行く方が楽だと思う
	// 「整数のN桁目がX」は文字列で処理したほうが楽
	for i := 0; i < int(math.Pow10(N)); i++ {
		n := strconv.Itoa(i)

		// 桁数がおかしい場合
		if len(n) != N {
			continue
		}

		ok := true

		// 条件を満たさないことがあれば false
		for j := 0; j < M; j++ {
			s, c := SC[j][0], SC[j][1]
			s--

			if int(n[s]-'0') != c {
				ok = false
			}
		}

		if ok {
			return i
		}
	}

	return -1

}

func TestAnswerABC157Cその2(t *testing.T) {
	tests := []struct {
		name string
		N, M int
		SC   [][]int
		want int
	}{
		{
			"入力例1",
			3, 3,
			[][]int{
				{1, 7},
				{3, 2},
				{1, 7},
			},
			702,
		},
		{
			"入力例2: 情報の矛盾",
			3, 2,
			[][]int{
				{2, 1},
				{2, 3},
			},
			-1,
		},
		{
			"入力例3: 上一桁が0はOUT",
			3, 1,
			[][]int{
				{1, 0},
			},
			-1,
		},
		{
			"答えが0",
			1, 1,
			[][]int{
				{1, 0},
			},
			0,
		},
		{
			"最大桁の情報がない場合は、0じゃなくて1",
			2, 1,
			[][]int{
				{2, 3},
			},
			13,
		},
		{
			"最大桁の情報がない場合は、0じゃなくて1だけど、103みたいになる",
			3, 1,
			[][]int{
				{3, 3},
			},
			103,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC157Cその2(tt.N, tt.M, tt.SC)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
