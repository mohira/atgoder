package topic_cumlative_sum

import (
	"testing"
)

// [ABC024B - 自動ドア](https://atcoder.jp/contests/abc024/tasks/abc024_b)
func AnswerABC024Bその2(N, T int, A []int) int {
	//
	// 2つの区間がかぶっている場合(必ず Ai<=Aj)
	//       b       b+T
	//       |--------|
	//   |--------|
	//   a       a+T
	//    *** *********

	var totalTime int
	A = append(A, A[N-1]+T+1) // 最後の時刻のために追加

	for i := 0; i < N; i++ {
		now := A[i]
		next := A[i+1]

		if now+T < next {
			totalTime += T
		} else {
			totalTime += next - now
		}
	}

	return totalTime
}

func TestAnswerABC024Bその2(t *testing.T) {
	tests := []struct {
		name string
		N, T int
		A    []int
		want int
	}{
		{"小さい例",
			3, 5,
			[]int{5, 15, 17},
			12},
		{"入力例1",
			5, 10,
			[]int{20, 100, 105, 217, 314},
			45},
		{"入力例2",
			10, 10,
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			19},
		{"入力例3",
			10, 100000,
			[]int{3, 31, 314, 3141, 31415, 314159, 400000, 410000, 500000, 777777},
			517253,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC024Bその2(tt.N, tt.T, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
