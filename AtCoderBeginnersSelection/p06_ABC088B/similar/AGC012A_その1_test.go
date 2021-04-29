package similar

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"testing"
)

// [AGC012A - AtCoder Group Contest](https://atcoder.jp/contests/agc012/tasks/agc012_a)
func AnswerAGC012Aその1(N int, A []int) int {
	var ans int
	//sort.Sort(sort.Reverse(sort.IntSlice(A)))
	//for i := 0; i < 2*N; i += 2 {
	//	ans += A[i+1]
	//}
	sort.Ints(A)
	for i := len(A) - 1; i > N; i -= 2 {
		ans += A[i-1]
	}

	return ans
}

func zeroInts(n int) []int {
	return make([]int, 3*n)
}

func TestAnswerAGC012Aその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want int
	}{
		{"入力例1", 2, []int{5, 2, 8, 5, 1, 5}, 10},
		{"入力例追加", 3, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 18},
		{"入力例2", 10, []int{1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000}, 10000000000},
		{"TLE 06.txt(原因は入力の受け取りだった)", 100_000, zeroInts(100_000), 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerAGC012Aその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}

// 以下、提出コード
// 試す場合は `AGC012_06.txt` を入力リダイレクトが便利
var scanner = bufio.NewScanner(os.Stdin)

func getI() int {
	scanner.Scan()
	i, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func getInts(N int) []int {
	ints := make([]int, N)
	for i := 0; i < N; i++ {
		ints[i] = getI()
	}
	return ints
}

func main() {
	scanner.Split(bufio.ScanWords)
	scanner.Buffer([]byte{}, math.MaxInt32)

	N := getI()

	A := getInts(N * 3)

	fmt.Println(AnswerAGC012Aその1(N, A))
}
