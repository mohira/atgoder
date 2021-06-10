package ABC_problemB

import (
	"testing"
)

// [ABC188B - Orthogonality](https://atcoder.jp/contests/abc188/tasks/abc188_b)
func AnswerABC188Bその1(N int, A, B []int) string {
	var product int
	for i := 0; i < N; i++ {
		a, b := A[i], B[i]
		product += a * b

	}

	if product == 0 {
		return "Yes"
	} else {
		return "No"
	}
}

func TestAnswerABC188Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		B    []int
		want string
	}{
		{"入力例1",
			2,
			[]int{-3, 6},
			[]int{4, 2},
			"Yes",
		},

		{"入力例2",
			2,
			[]int{4, 5},
			[]int{-1, -3},
			"No",
		},

		{"入力例3",
			3,
			[]int{1, 3, 5},
			[]int{3, -6, 3},
			"Yes",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC188Bその1(tt.N, tt.A, tt.B)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
