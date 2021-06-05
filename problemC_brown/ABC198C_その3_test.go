package problemC_brown

import (
	"testing"
)

// [ABC198C - Compass Walking](https://atcoder.jp/contests/abc198/tasks/abc198_c)
func AnswerABC198Cその3(R, X, Y int) int {
	// 整数の世界でやるパターン: https://youtu.be/wHgVlM0UVO4?t=1333
	// 「N歩でいけるか？」を考える
	// N歩進めたときに、目標地点を通り過ぎるかどうか

	d := X*X + Y*Y

	var N = 1
	for {
		if R*R*N*N >= d {
			break
		}
		N++
	}

	if N == 1 {
		if R*R != d {
			N = 2
		}
	}

	return N
}

func TestAnswerABC198Cその3(t *testing.T) {
	tests := []struct {
		name    string
		R, X, Y int
		want    int
	}{
		{"入力例1", 5, 15, 0, 3},
		{"入力例2", 5, 11, 0, 3},
		{"入力例3", 3, 4, 4, 2},
		{"", 95044, 28016, 39332, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC198Cその3(tt.R, tt.X, tt.Y)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
