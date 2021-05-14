package p08_ABC085C

import (
	"math"
	"testing"
)

// [ABC183B - Billiards](https://atcoder.jp/contests/abc183/tasks/abc183_b)
func AnswerABC183Bその1(Sx, Sy, Gx, Gy int) float64 {
	// tanS と tanG が一致すればいい
	// tanの正負は一致するのでそのまま割ればOK
	//      Sx < Gx => tanはどちらも正
	// Gx < Sx      => tanはどちらも負
	return float64(Sy*Gx+Gy*Sx) / float64(Sy+Gy)
}

func TestAnswerABC183Bその1(t *testing.T) {
	tests := []struct {
		name           string
		Sx, Sy, Gx, Gy int
		want           float64
	}{
		{"入力例1", 1, 1, 7, 2, 3.0000000000},
		{"入力例2", 1, 1, 3, 2, 1.6666666667},
		{"入力例3", -9, 99, -999, 9999, -18.7058823529},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC183Bその1(tt.Sx, tt.Sy, tt.Gx, tt.Gy)

			if math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
