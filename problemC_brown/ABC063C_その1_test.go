package problemC_brown

import (
	"atgoder/lib"
	"testing"
)

// [ABC063C - Bugged](https://atcoder.jp/contests/abc063/tasks/abc063_c)
func AnswerABC063Cその1(N int, S []int) int {
	sum := lib.SumInts(S)

	// 総和が10の倍数でないなら、そのまま答え
	if sum%10 != 0 {
		return sum
	}

	// 総和が10の倍数なら、10の倍数を崩すような最小の値を1回だけ引くのがベスト
	max := 0
	for _, s := range S {
		tmp := sum - s

		if tmp%10 != 0 && max < tmp {
			max = tmp
		}
	}

	return max
}

func TestAnswerABC063Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		S    []int
		want int
	}{
		{"入力例1", 3, []int{5, 10, 15}, 25},
		{"入力例2", 3, []int{10, 10, 15}, 35},
		{"入力例3", 3, []int{10, 20, 30}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC063Cその1(tt.N, tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
