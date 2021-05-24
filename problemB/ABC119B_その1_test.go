package problemB

import (
	"math"
	"testing"
)

// [ABC119B - Can you solve this?](https://atcoder.jp/contests/abc119/tasks/abc119_b)
func AnswerABC119Bその1(N int, X []float64, U []string) float64 {
	const BTC float64 = 380000.0

	var totalYen float64
	for i := 0; i < N; i++ {
		x, u := X[i], U[i]

		if u == "BTC" {
			totalYen += x * BTC
		} else {
			totalYen += x * 1
		}
	}

	return totalYen
}

func TestAnswerABC119Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		X    []float64
		U    []string
		want float64
	}{
		{
			"入力例1",
			2,
			[]float64{10000, 0.10000000},
			[]string{"JPY", "BTC"},
			48000.0,
		},
		{
			"入力例2",
			3,
			[]float64{100000000, 100.00000000, 0.00000001},
			[]string{"JPY", "BTC", "BTC"},
			138000000.0038,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC119Bその1(tt.N, tt.X, tt.U)

			if math.Abs(got-tt.want) > 1e-5 {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
