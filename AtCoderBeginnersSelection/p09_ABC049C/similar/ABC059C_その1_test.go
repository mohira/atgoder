package p08_ABC049C

import (
	"fmt"
	"os"
	"testing"
)

// [ABC059C - Sequence](https://atcoder.jp/contests/abc059/tasks/arc072_a)
func AnswerABC059Cその1(N int, A []int) int {
	odd := count初項を正とする場合の操作回数(A)
	even := count初項を負とする場合の操作回数(A)

	return min(odd, even)
}

func count初項を負とする場合の操作回数(A []int) int {
	B := A

	var sum int
	var opCount int

	for i, a := range A {
		pprint(B)

		sum += a
		// 初項 < 0  としているので、奇数項までの 累積和<0 でないといけない
		// 言い換えると、 累積和 >= 0 なら、-1操作を必要な数だけしないといけない
		if i%2 == 0 && sum >= 0 {
			// 累積和が -1 になるように操作する
			minusOperation := abs(sum) + 1
			sum -= minusOperation
			opCount += minusOperation

			B = update(B, i, a)
		}

		// 初項 < 0  としているので、偶数項までの 累積和>0 でないといけない
		// 言い換えると、 累積和 <= 0 なら、+1操作を必要な数だけしないといけない
		if i%2 == 1 && sum <= 0 {
			// 累積和が +1 になるように操作する
			plusOperation := abs(sum) + 1
			sum += plusOperation
			opCount += plusOperation

			B = update(B, i, a)
		}
	}

	return opCount
}

func count初項を正とする場合の操作回数(A []int) int {
	// ステップごとの操作をわかりやすくするためのコピーを用意する
	B := A

	var sum int
	var opCount int

	for i, a := range A {
		pprint(B)

		sum += a

		// 初項 > 0 としているので、奇数項までの 累積和>0 でないといけない
		// 言い換えると、 累積和 <= 0 なら、+1操作を必要な数だけしないといけない
		if i%2 == 0 && sum <= 0 {
			// 累積和が +1 になるように操作する
			plusOperationCount := abs(sum) + 1
			sum += plusOperationCount
			opCount += plusOperationCount

			B = update(B, i, a)
		}

		// 初項 > 0 としているので、偶数項までの 累積和<0 でないといけない
		// 言い換えると、 累積和 >= 0 なら、-1操作を必要な数だけしないといけない
		if i%2 == 1 && sum >= 0 {
			// 累積和が -1 になるように操作する
			minusOperationCount := abs(sum) + 1

			sum -= minusOperationCount
			opCount += minusOperationCount

			B = update(B, i, a)
		}
	}

	return opCount
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}

}

func update(B []int, i int, a int) []int {
	B[i] = a
	return B
}

// デバッグ用の出力関数
func pprint(A []int) {
	// index
	fmt.Fprint(os.Stderr, " i | ")
	for i, _ := range A {
		fmt.Fprintf(os.Stderr, " %d ", i)
	}
	fmt.Fprintln(os.Stderr)
	fmt.Fprint(os.Stderr, "ai | ")

	// ai
	for _, a := range A {
		fmt.Fprintf(os.Stderr, "%+d ", a)
	}
	fmt.Fprintln(os.Stderr)

	// 累積和
	var sum int
	fmt.Fprint(os.Stderr, "Ri | ")
	for _, a := range A {
		sum += a
		fmt.Fprintf(os.Stderr, "%+d ", sum)
	}
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "======================")
}

func TestAnswerABC059Cその2(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want int
	}{
		{"入力例1: 初項 > 0 が正解", 4, []int{+1, -3, +1, 0}, 4},
		{"入力例2: 最初から条件をみたしている", 5, []int{3, -6, 4, -5, 7}, 0},
		{"入力例3: 初項 < 0 が正解", 5, []int{-1, 4, 3, 2, -5, 4}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC059Cその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
