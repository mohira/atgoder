package problemB

import (
	"testing"
)

// [ABC078B - ISU](https://atcoder.jp/contests/abc078/tasks/abc078_b)
func AnswerABC078Bその1(X, Y, Z int) int {
	X -= Z
	count := 0
	for {
		if !(Z <= X) {
			break
		}
		X -= Y

		if !(Z <= X) {
			break
		}
		X -= Z

		count++
	}
	return count
}

func TestAnswerABC078Bその1(t *testing.T) {
	tests := []struct {
		name    string
		X, Y, Z int
		want    int
	}{
		{"入力例1", 13, 3, 1, 3},
		{"入力例2", 12, 3, 1, 2},
		{"入力例3", 100000, 1, 1, 49999},
		{"入力例4", 64146, 123, 456, 110},
		{"入力例5", 64145, 123, 456, 109},
		{"最小ケース", 3, 1, 1, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC078Bその1(tt.X, tt.Y, tt.Z)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
