package p08_ABC049C

import (
	"sort"
	"testing"
)

// [AGC011A - Airport Bus](https://atcoder.jp/contests/agc011/tasks/agc011_a)
func AnswerAGC011Aその1(N, C, K int, T []int) int {
	// 時刻と乗車人数を並行してカウントするのが難しかったので、シミュレーション的に実装している
	// 時刻はbus[0]で、乗車人数はlen(bus)で管理している
	sort.Ints(T)

	// まずは1人目を乗せる
	var bus = []int{T[0]}
	var busCount = 1

	// 1人目を乗せているので 2人目(i=1)からスタート
	for i := 1; i < N; i++ {
		t := T[i]

		allSeatsAreAvailable := len(bus) == 0
		allSeatsAreFull := C == len(bus)
		hasAvailableSeats := len(bus) <= C-1

		switch {
		case allSeatsAreAvailable:
			// 誰も乗車していないなら、誰でも乗れる
			bus = append(bus, t)

		case allSeatsAreFull:
			// バスの座席容量がないなら、新たなバスに乗るしかない
			bus = []int{t}
			busCount++

		case hasAvailableSeats:
			// バスの座席容量が余っているとき
			theBusArrivedTime := bus[0] // そのバスに最初に乗車した人の時刻が基準
			willNotAngry := t <= theBusArrivedTime+K

			if willNotAngry {
				// 時刻的にOKなら乗る
				bus = append(bus, t)
			} else {
				// 時刻的にNGなら、新たなバスに乗る
				bus = []int{t}
				busCount++
			}
		}
	}

	return busCount
}

func TestAnswerAGC011Aその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		C    int
		K    int
		T    []int
		want int
	}{
		{"入力例1", 5, 3, 5, []int{1, 2, 3, 6, 12}, 3},
		{"入力例2", 6, 3, 3, []int{7, 6, 2, 8, 10, 6}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerAGC011Aその1(tt.N, tt.C, tt.K, tt.T)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
