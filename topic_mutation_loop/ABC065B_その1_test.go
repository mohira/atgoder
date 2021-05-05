package topic_grid

import (
	"testing"
)

// [ABC065B - Trained?](https://atcoder.jp/contests/abc065/tasks/abc065_b)
func AnswerABC065Bその1(N int, A []int) int {
	/// 素直に光っているボタンを押すことを繰り返す

	// {押すボタン: 光るボタン} のマップ
	routeMap := make(map[int]int, N)
	for i := 0; i < N; i++ {
		routeMap[i+1] = A[i]
	}

	// 光るボタン用のバケット
	lightOnBucket := make(map[int]int)

	nowLight := 1
	opCount := 0
	for {
		nextLight := routeMap[nowLight]
		opCount++

		// ボタン2が光ったらおしまい
		if nextLight == 2 {
			return opCount
		}

		// 光るボタンの周期に入ったら抜ける
		if _, ok := lightOnBucket[nextLight]; ok {
			break
		}

		lightOnBucket[nextLight]++
		nowLight = nextLight
	}

	return -1
}

func TestAnswerABC065Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want int
	}{
		{"入力例1", 3, []int{3, 1, 2}, 2},
		{"入力例2", 4, []int{3, 4, 1, 2}, -1},
		{"入力例3", 5, []int{3, 3, 4, 2, 4}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC065Bその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
