package topic_bit_brute_force

import (
	"testing"
)

// [部分和問題]
// N個の正の整数a0,a1,...,aN－1と正の整数Wが与えられます。
// a0,a1,...,aN－1の中から何個かの整数を選んで総和をWとすることができるかどうかを判定してください。
// 大槻兼資. 問題解決力を鍛える！アルゴリズムとデータ構造 (Japanese Edition) (Kindle の位置No.911-916). Kindle 版.

func Answer部分和(N int, W int, A []int) string {
	var can bool

	for bit := 0; bit < (1 << N); bit++ {
		sum := 0
		for i := 0; i < N; i++ {
			if bit&(1<<i) > 0 {
				sum += A[i]
			}
		}
		if sum == W {
			can = true
		}
	}

	if can {
		return "YES"
	} else {
		return "NO"
	}
}

func TestAnswer部分和(t *testing.T) {
	tests := []struct {
		name string
		N    int
		W    int
		A    []int
		want string
	}{
		{"入力例1", 3, 5, []int{1, 2, 3}, "YES"},
		{"入力例2", 5, 10, []int{1, 2, 4, 5, 11}, "YES"},
		{"入力例3", 4, 10, []int{1, 5, 8, 11}, "NO"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Answer部分和(tt.N, tt.W, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
