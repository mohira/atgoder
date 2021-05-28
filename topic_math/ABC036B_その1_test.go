package topic_math

import (
	"bytes"
	"strings"
	"testing"
)

// [ABC036B - 回転](https://atcoder.jp/contests/abc036/tasks/abc036_b)
func AnswerABC036Bその1(N int, S []string) string {
	newS := make([][]string, N)
	for i := 0; i < N; i++ {
		newS[i] = make([]string, N)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			newX, newY := ToMinus90degreesRotation(N, i, j)

			newS[newX][newY] = string(S[i][j])
		}
	}

	var out bytes.Buffer
	for _, s := range newS {
		out.WriteString(strings.Join(s, "") + "\n")
	}
	return out.String()
}

// -90度回転する座標を求める(マイナス座標の補正をしている)
func ToMinus90degreesRotation(N int, x int, y int) (int, int) {
	// -90度回転行列
	//  0 1
	// -1 0
	a11, a12 := 0, 1
	a21, a22 := -1, 0

	newX := a11*x + a12*y
	newY := a21*x + a22*y + N - 1 // +N はマイナス座標補正 // -1 はidx補正

	return newX, newY
}

func TestAnswerABC036Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		S    []string
		want string
	}{
		{
			"入力例1",
			4,
			[]string{
				"ooxx",
				"xoox",
				"xxxx",
				"xxxx",
			},
			"xxxo\nxxoo\nxxox\nxxxx\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC036Bその1(tt.N, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
