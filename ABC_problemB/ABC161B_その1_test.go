package ABC_problemB

import (
	"testing"
)

// [ABC161B - Popular Vote](https://atcoder.jp/contests/abc161/tasks/abc161_b)
func AnswerABC161Bその1(N, M int, A []int) string {
	var totalVotes int
	for _, a := range A {
		totalVotes += a
	}

	var count int
	for _, vote := range A {
		if totalVotes <= vote*(4*M) {
			count++
		}
	}

	if M <= count {
		return "Yes"
	} else {
		return "No"
	}
}

func TestAnswerABC161Bその1(t *testing.T) {
	tests := []struct {
		name string
		N, M int
		A    []int
		want string
	}{
		{"入力例1",
			4, 1,
			[]int{5, 4, 2, 1},
			"Yes",
		},
		{"入力例2",
			3, 2,
			[]int{380, 19, 1},
			"No",
		},
		{"入力例3",
			12, 3,
			[]int{4, 56, 78, 901, 2, 345, 67, 890, 123, 45, 6, 789},
			"Yes",
		},
		{"除算をすると精度が問題になるケース: hand_06",
			10, 3,
			[]int{100, 99, 1, 2, 3, 4, 5, 10, 19, 22},
			"No",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC161Bその1(tt.N, tt.M, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
