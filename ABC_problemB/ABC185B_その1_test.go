package ABC_problemB

import (
	"atgoder/lib"
	"testing"
)

// [ABC185B - Smartphone Addiction](https://atcoder.jp/contests/abc185/asks/abc185_b)
func AnswerABC185Bその1(N, M, T int, AB [][]int) string {
	battery := N

	出発時刻 := 0

	for i := 0; i < M; i++ {
		到着時刻, 出発予定時刻 := AB[i][0], AB[i][1]

		// 歩行時間分だけバッテリー消費
		歩行時間 := 到着時刻 - 出発時刻
		battery -= 歩行時間

		if battery <= 0 {
			return "No"
		}

		// 滞在時間分だけバッテリー回復(ただし、上限があることに注意)
		滞在時間 := 出発予定時刻 - 到着時刻
		battery = lib.Min(N, battery+滞在時間)

		// 入れ替え
		出発時刻 = 出発予定時刻
	}

	// 最後のカフェから自宅への移動によりバッテリー消費
	battery -= T - AB[M-1][1]

	if battery <= 0 {
		return "No"
	}

	return "Yes"
}

func TestAnswerABC185Bその1(t *testing.T) {
	tests := []struct {
		name    string
		N, M, T int
		AB      [][]int
		want    string
	}{
		{"入力例1",
			10, 2, 20,
			[][]int{
				{9, 11},
				{13, 17},
			},
			"Yes",
		},

		{"入力例2",
			10, 2, 20,
			[][]int{
				{9, 11},
				{13, 16},
			},
			"No",
		},
		{"入力例3",
			15, 3, 30,
			[][]int{
				{5, 8},
				{15, 17},
				{24, 27},
			},
			"Yes",
		},
		{"入力例4",
			20, 1, 30,
			[][]int{
				{20, 29},
			},
			"No",
		},
		{"バッテリー上限を超えての回復はしないケース",
			20, 1, 30,
			[][]int{
				{1, 10},
			},
			"No",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC185Bその1(tt.N, tt.M, tt.T, tt.AB)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
