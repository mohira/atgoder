package problemB

import (
	"fmt"
	"testing"
)

// [ABC093B - Small and Large Integers](https://atcoder.jp/contests/abc093/tasks/abc093_b)
func AnswerABC093Bその1(A, B, K int) string {
	// ex: A=4, B=8, K=3
	// A   A+K A+K-1
	// 4    5    6    7    8
	// |    |    |
	//           |    |    |
	//         B-K+1       B
	var ans string
	for i := A; i <= min(A+K-1, B); i++ {
		ans += fmt.Sprintf("%d\n", i)
	}

	for i := max(A+K, B-K+1); i <= B; i++ {
		ans += fmt.Sprintf("%d\n", i)
	}

	return ans
}

func TestAnswerABC093Bその1(t *testing.T) {
	tests := []struct {
		name    string
		A, B, K int
		want    string
	}{
		{
			"入力例1",
			3, 8, 2,
			"3\n4\n7\n8\n",
		},
		{
			"入力例2",
			4, 8, 3,
			"4\n5\n6\n7\n8\n",
		},
		{
			"入力例3",
			2, 9, 100,
			"2\n3\n4\n5\n6\n7\n8\n9\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC093Bその1(tt.A, tt.B, tt.K)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
