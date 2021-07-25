package topic_bit_brute_force

import (
	"strconv"
	"testing"
)

// [ABC079C - Train Ticket](https://atcoder.jp/contests/abc079/tasks/abc079_c)
func AnswerABC079Cその1(ABCD string) string {
	// 扱いやすくするために整数スライスに変換する
	var A = make([]int, 4)
	for i, s := range ABCD {
		a, _ := strconv.Atoi(string(s))
		A[i] = a
	}

	var N = 3
	var op string

	for bit := 0; bit < (1 << N); bit++ {
		a0 := A[0]
		sum := a0
		statement := strconv.Itoa(a0)

		for i := 0; i < N; i++ {
			a := A[i+1]

			if bit&(1<<i) > 0 {
				op = "+"
				sum += a
			} else {
				op = "-"
				sum -= a
			}
			statement += op + strconv.Itoa(a)
		}

		if sum == 7 {
			return statement + "=7"
		}
	}

	return "" // 必ず 7 がつくれるのでここには到達しない
}

func TestAnswerABC079Cその1(t *testing.T) {
	tests := []struct {
		name string
		ABCD string
		want string
	}{
		{"入力例1", "1222", "1+2+2+2=7"},
		{"入力例2", "0290", "0-2+9-0=7"}, // "0−2+9+0=7" でもOK
		{"入力例3", "3242", "3+2+4-2=7"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC079Cその1(tt.ABCD)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
