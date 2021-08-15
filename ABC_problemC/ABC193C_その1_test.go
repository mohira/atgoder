package ABC_problemC

import (
	"bytes"
	"strconv"
	"testing"
)

// [ABC157C - Guess The Number](https://atcoder.jp/contests/abc157/tasks/abc157_c)
func AnswerABC157Cその1(N, M int, SC [][]int) string {
	// 整数が存在しない場合を攻略する
	mujun := make(map[int]int)
	for i := 0; i < M; i++ {
		s, c := SC[i][0], SC[i][1]
		// 2桁以上の整数の上1ケタ0は矛盾
		if N >= 2 && s == 1 && c == 0 {
			return "-1"
		}

		// 同じ桁数に対して別の情報が来た場合は矛盾
		if _, ok := mujun[s]; ok && mujun[s] != c {
			return "-1"
		} else {
			mujun[s] = c
		}
	}

	// 整数バケットで管理すればインデックス通りに読めばよいのでわかりやすい
	bucket := make([]int, N)

	for i := 0; i < M; i++ {
		s, c := SC[i][0], SC[i][1]
		s--
		bucket[s] = c
	}

	// 2桁以上の整数の場合で、上1桁が決まってない場合は1にセットするのが最小
	if N >= 2 && bucket[0] == 0 {
		bucket[0] = 1
	}

	var out bytes.Buffer
	for _, v := range bucket {
		out.WriteString(strconv.Itoa(v))
	}
	return out.String()
}

func TestAnswerABC157Cその1(t *testing.T) {
	tests := []struct {
		name string
		N, M int
		SC   [][]int
		want string
	}{
		{
			"入力例1",
			3, 3,
			[][]int{
				{1, 7},
				{3, 2},
				{1, 7},
			},
			"702",
		},
		{
			"入力例2: 情報の矛盾",
			3, 2,
			[][]int{
				{2, 1},
				{2, 3},
			},
			"-1",
		},
		{
			"入力例3: 上一桁が0はOUT",
			3, 1,
			[][]int{
				{1, 0},
			},
			"-1",
		},
		{
			"答えが0",
			1, 1,
			[][]int{
				{1, 0},
			},
			"0",
		},
		{
			"最大桁の情報がない場合は、0じゃなくて1",
			2, 1,
			[][]int{
				{2, 3},
			},
			"13",
		},
		{
			"最大桁の情報がない場合は、0じゃなくて1だけど、103みたいになる",
			3, 1,
			[][]int{
				{3, 3},
			},
			"103",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC157Cその1(tt.N, tt.M, tt.SC)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
