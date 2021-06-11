package AGC_problemA

import (
	"atgoder/lib"
	"log"
	"strconv"
	"testing"
)

// [AGC021 - Digit Sum 2](https://atcoder.jp/contests/agc021/tasks/agc021_a)
func AnswerAGC021その1(N int) int {
	s := strconv.Itoa(N)

	// 最上位ケタ以外を9、最上位ケタを1減らす
	// 314→299
	// 100→ 99
	keta := len(s)
	head, err := strconv.Atoi(string(s[0]))
	if err != nil {
		log.Fatal(err)
	}
	candidate := 9*(keta-1) + (head - 1)

	// 1桁ケースと、全て9の場合をケアするために、もともとのNと比較
	//   9→  9
	// 999→899
	digitSum := 0
	for _, c := range s {
		digitSum += int(c - '0')
	}

	return lib.Max(digitSum, candidate)
}

func TestAnswerAGC021その1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{"入力例1", 100, 18},
		{"入力例2", 9995, 35},
		{"入力例3", 3141592653589793, 137},

		{"最上位ケタ以外が9", 19, 10},
		{"最上位ケタ以外が9", 199, 19},

		{"最上位ケタが9", 9, 9},
		{"最上位ケタが9", 99, 18},

		{"1ケタ", 9, 9},
		{"1ケタ", 8, 8},
		{"1ケタ", 1, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerAGC021その1(tt.N)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
