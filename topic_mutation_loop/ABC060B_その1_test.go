package topic_grid

import (
	"testing"
)

// [ABC060B - Choose Integers](https://atcoder.jp/contests/abc060/tasks/abc060_b)
func AnswerABC060Bその1(A, B, C int) string {
	// Aの倍数の総和はどう選んでも、Aの倍数なので、結局1つだけAの倍数を選べばいい
	// N mod B は 周期的になるので、1つ分の周期をバケット法で探す
	// バケット内にCがあればYES
	var amariBucket = make(map[int]int)

	var i = 1
	for {
		amari := (A * i) % B

		if _, ok := amariBucket[amari]; ok {
			break
		}

		amariBucket[amari]++
		i++
	}

	if _, ok := amariBucket[C]; ok {
		return "YES"
	} else {
		return "NO"
	}
}

func TestAnswerABC060Bその1(t *testing.T) {
	tests := []struct {
		name string
		A    int
		B    int
		C    int
		want string
	}{
		{"入力例1", 7, 5, 1, "YES"},
		{"入力例2", 2, 2, 1, "NO"},
		{"入力例3", 1, 100, 97, "YES"},
		{"入力例4", 40, 98, 58, "YES"},
		{"入力例5", 77, 42, 36, "NO"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC060Bその1(tt.A, tt.B, tt.C)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
