package p06_ABC071B

import (
	"fmt"
	"os"
	"testing"
)

// [ABC047B - Counting Roads](https://atcoder.jp/contests/abc047/tasks/abc047_b)
func AnswerABC047Bその1(W, H, N int, XYA [][3]int) int {
	var bucket = make([][]int, H)
	for i := 0; i < H; i++ {
		bucket[i] = make([]int, W)
	}
	pprint(bucket)

	for _, xya := range XYA {
		x, y, a := xya[0], xya[1], xya[2]
		if a == 1 {
			bucket = op1(bucket, x, y, W, H)
		}

		if a == 2 {
			bucket = op2(bucket, x, y, W, H)
		}

		if a == 3 {
			bucket = op3(bucket, x, y, W, H)
		}

		if a == 4 {
			bucket = op4(bucket, x, y, W, H)
		}
		pprint(bucket)
	}

	return countZeroArea(bucket)
}

func countZeroArea(bucket [][]int) int {
	// 0(==塗りつぶされていない場所)の数を数えればよい
	var area int

	for i := 0; i < len(bucket); i++ {
		for j := 0; j < len(bucket[i]); j++ {
			if bucket[i][j] == 0 {
				area++
			}

		}
	}
	return area
}

func op1(bucket [][]int, x int, y int, w int, h int) [][]int {
	for i := 0; i < x; i++ {
		for j := 0; j < h; j++ {
			bucket[j][i]++
		}
	}
	return bucket
}

func op2(bucket [][]int, x int, y int, w int, h int) [][]int {
	for i := x; i < w; i++ {
		for j := 0; j < h; j++ {
			bucket[j][i]++
		}
	}
	return bucket

}

func op3(bucket [][]int, x int, y int, w int, h int) [][]int {
	for i := 0; i < w; i++ {
		for j := 0; j < y; j++ {
			bucket[j][i]++
		}
	}
	return bucket

}

func op4(bucket [][]int, x int, y int, w int, h int) [][]int {
	for i := 0; i < w; i++ {
		for j := y; j < h; j++ {
			bucket[j][i]++
		}
	}
	return bucket
}

// デバッグ用
func pprint(bucket [][]int) {
	for i := len(bucket) - 1; i > 0; i-- {
		fmt.Fprintf(os.Stderr, "%02d %v\n", i, bucket[i])
	}
	fmt.Fprint(os.Stderr, "    ")
	for i := 1; i <= len(bucket[0]); i++ {
		fmt.Fprintf(os.Stderr, "%d ", i)
	}

	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr)
}

func TestAnswerABC047Bその2(t *testing.T) {
	tests := []struct {
		name string
		W    int
		H    int
		N    int
		XYA  [][3]int
		want int
	}{
		{"入力例1",
			5,
			4,
			2,
			[][3]int{
				{2, 1, 1},
				{3, 3, 4},
			},
			9,
		},
		{"入力例1の条件変え",
			5,
			4,
			2,
			[][3]int{
				{2, 1, 2},
				{3, 3, 3},
			},
			2,
		},
		{"入力例2",
			5,
			4,
			3,
			[][3]int{
				{2, 1, 1},
				{3, 3, 4},
				{1, 4, 2},
			},
			0,
		},
		{"入力例3",
			10,
			10,
			5,
			[][3]int{
				{1, 6, 1},
				{4, 1, 3},
				{6, 9, 4},
				{9, 4, 2},
				{3, 1, 3},
			},
			64,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC047Bその1(tt.W, tt.H, tt.N, tt.XYA)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
