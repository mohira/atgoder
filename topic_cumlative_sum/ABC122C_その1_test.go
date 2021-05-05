package topic_cumlative_sum

import (
	"fmt"
	"testing"
)

// [ABC122C - GeT AC](https://atcoder.jp/contests/abc122/tasks/abc122_c)
func AnswerABC122Cその1(N, Q int, S string, LR [][2]int) string {
	// MEMO: 最後の文字列結合があるとTLEになるので注意
	//       (テストしやすくするために出力コードは書かないポリシーだが、ACするためにはあきらめないといけない)

	//      i  0 1 2 3 4 5 6 7 8
	//           A C A C T A C G
	// Q(3,7)        A C T A C
	// Q(2,3)      C A
	// Q(1,8)    A C A C T A C G
	//         -----------------
	// sum     0 0 1 1 2 2 2 3 3

	cumSum := ACの出現回数の累積和(N, S)

	var ans string
	for i := 0; i < Q; i++ {
		li, ri := LR[i][0], LR[i][1]

		// li文字目からri文字目までの "AC" 出現回数
		countAC := cumSum[ri] - cumSum[li]
		ans += fmt.Sprintf("%d\n", countAC)
	}

	return ans
}

func ACの出現回数の累積和(N int, S string) []int {
	cumSum := make([]int, 1, len(S)+1)

	var prev, now rune
	var count int

	for i := 0; i < N; i++ {
		// strings.Count を使うと中でforループするのでO(N)になってしまうから、自力でカウントしている
		prev, now = now, rune(S[i])
		if prev == 'A' && now == 'C' {
			count++
		}
		cumSum = append(cumSum, count)
	}

	return cumSum
}

func main() {
	// ACするコード
	// 最大で 1200msくらい
	// 入出力まわりでの高速化もできそう
	//
	// https://atcoder.jp/contests/abc122/submissions?f.Task=abc122_c&f.LanguageName=Go&f.Status=AC も参考にするとよさげ
	var N, Q int
	var S string
	_, _ = fmt.Scan(&N, &Q, &S)

	cumSum := ACの出現回数の累積和(N, S)

	var li, ri int

	for i := 0; i < Q; i++ {
		_, _ = fmt.Scan(&li, &ri)

		// li文字目からri文字目までの "AC" 出現回数
		countAC := cumSum[ri] - cumSum[li]
		fmt.Println(countAC)
	}
}

func TestAnswerABC122Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		Q    int
		S    string
		LR   [][2]int
		want string
	}{
		{"入力例1",
			8,
			3,
			"ACACTACG",
			[][2]int{
				{3, 7},
				{2, 3},
				{1, 8},
			},
			"2\n0\n3\n",
		},
		{"",
			4,
			4,
			"ACAC",
			[][2]int{
				{1, 2},
				{1, 3},
				{1, 4},
				{2, 3},
			},
			"1\n1\n2\n0\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC122Cその1(tt.N, tt.Q, tt.S, tt.LR)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
