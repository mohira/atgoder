package similar

import (
	"testing"
)

// [ABC175B - Good Distance](https://atcoder.jp/contests/abc175/tasks/abc175_b)
func AnswerABC175Bその1(N int, L []int) int {
	var cnt int

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			for k := j + 1; k < N; k++ {
				Li, Lj, Lk := L[i], L[j], L[k]

				if isAllDifferent(Li, Lj, Lk) && canExistTriangle(Li, Lj, Lk) {
					cnt++
				}
			}
		}
	}
	return cnt
}

func isAllDifferent(a int, b int, c int) bool {
	return (a != b) && (b != c) && (c != a)
}

// [三角形の成立条件とその証明 | 高校数学の美しい物語](https://manabitimes.jp/math/862)
func canExistTriangle(a int, b int, c int) bool {
	return a+b > c && b+c > a && c+a > b

	// |b-c| < a < ∣b−c∣ でもOK (けど、floatが入るのが嫌だなあ)
	// var x, y, z = float64(a), float64(b), float64(c)
	// return math.Abs(y-z) < x && x < math.Abs(y+z)
}

func TestAnswerABC175Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		L    []int
		want int
	}{
		{"入力例1", 5, []int{4, 4, 9, 7, 5}, 5},
		{"入力例2", 6, []int{4, 5, 4, 3, 3, 5}, 8},
		{"入力例3", 10, []int{9, 4, 6, 1, 9, 6, 10, 6, 6, 8}, 39},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC175Bその1(tt.N, tt.L)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}

		})

	}
}
