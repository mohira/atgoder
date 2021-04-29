package p06_ABC085B

import (
	"testing"
)

// [ABC085B - Kagami Mochi](https://atcoder.jp/contests/abc085/tasks/abc085_b)
func AnswerABC085Bその2(N int, D []int) int {
	// バケット法
	// [AtCoder に登録したら次にやること ～ これだけ解けば十分闘える！過去問精選 10 問 ～ - Qiita](https://qiita.com/drken/items/fd4e5e3630d0f5859067#%E7%AC%AC-7-%E5%95%8F--abc-085-b---kagami-mochi-200-%E7%82%B9)
	var ans int
	const (
		nMin       = 1
		nMax       = 100
		bucketSize = nMax + 1
	)

	//      D = [2, 3, 5, 3] なら
	// bucket = [0 1 2 0 1 0 0 0 0 ... ] になる
	// indexが数にそのまま対応する
	// バケットサイズは
	var bucket [bucketSize]int
	for i := 0; i < N; i++ {
		bucket[D[i]]++
	}

	// 1 <= N <= 100 なので、その範囲だけ探索すればOK
	for i := nMin; i <= nMax; i++ {
		if 0 < bucket[i] {
			ans++
		}
	}

	return ans
}

func TestAnswerABC085Bその2(t *testing.T) {
	tests := []struct {
		name string
		N    int
		D    []int
		want int
	}{
		{"入力例1", 4, []int{10, 8, 8, 6}, 3},
		{"入力例2", 3, []int{15, 15, 15}, 1},
		{"入力例3", 7, []int{50, 30, 50, 100, 50, 80, 30}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC085Bその2(tt.N, tt.D)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
