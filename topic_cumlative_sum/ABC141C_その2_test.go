package topic_cumlative_sum

import (
	"testing"
)

// [ABC141C - Attack Survival](https://atcoder.jp/contests/abc141/tasks/abc141_c)
func AnswerABC141Cその2(N, K, Q int, A []int) string {
	// 正解者以外 -1点 ではなく、最初からマイナスしておいて、正解者のみ +1点 の実装が楽

	points := make([]int, N)
	for i := 0; i < N; i++ {
		// 1問も正解していない状態は、 初期値 - 問題数
		points[i] = K - Q
	}

	for _, q := range A {
		q--
		points[q]++
	}

	var ans string

	for _, point := range points {
		if point <= 0 {
			ans += "No\n"
		} else {
			ans += "Yes\n"
		}
	}
	return ans
}

func TestAnswerABC141Cその2(t *testing.T) {
	tests := []struct {
		name    string
		N, K, Q int
		A       []int
		want    string
	}{
		{"入力例1",
			6, 3, 4,
			[]int{3, 1, 3, 2},
			"No\nNo\nYes\nNo\nNo\nNo\n",
		},

		{"右端の人が勝利した場合(index注意)",
			3, 2, 2,
			[]int{3, 3},
			"No\nNo\nYes\n",
		},

		{"入力例2",
			6, 5, 4,
			[]int{3, 1, 3, 2},
			"Yes\nYes\nYes\nYes\nYes\nYes\n",
		},

		{"入力例3",
			10, 13, 15,
			[]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9},
			"No\nNo\nNo\nNo\nYes\nNo\nNo\nNo\nYes\nNo\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC141Cその2(tt.N, tt.K, tt.Q, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
