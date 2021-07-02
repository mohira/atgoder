package p06_ABC071B

import (
	"sort"
	"testing"
)

type Ball struct {
	number int
	freq   int
}

type Balls []Ball

func (bs Balls) Len() int {
	return len(bs)
}

func (bs Balls) Less(i, j int) bool {
	if bs[i].freq < bs[j].freq {
		return true
	} else {
		return false
	}
}

func (bs Balls) Swap(i, j int) {
	bs[i], bs[j] = bs[j], bs[i]
}

// [ABC081C - Not so Diverse](https://atcoder.jp/contests/abc081/tasks/arc086_a)
func AnswerABC081Cその2(N, K int, A []int) int {
	bucket := make(map[int]int)
	for _, a := range A {
		bucket[a]++
	}

	balls := make(Balls, 0, len(bucket))
	for color, freq := range bucket {
		ball := Ball{color, freq}
		balls = append(balls, ball)
	}
	sort.Stable(balls)

	var ans int

	for i := 0; i < len(balls)-K; i++ {
		ans += balls[i].freq
	}

	return ans
}

func TestAnswerABC081Cその2(t *testing.T) {
	tests := []struct {
		name string
		N    int
		K    int
		A    []int
		want int
	}{
		{"おれおれケース: 整数順と頻度順が一致しないパターン", 6, 2, []int{1, 1, 2, 5, 5, 5}, 1},
		{"入力例1", 5, 2, []int{1, 1, 2, 2, 5}, 1},
		{"入力例2", 4, 4, []int{1, 1, 2, 2}, 0},
		{"入力例3", 10, 3, []int{5, 1, 3, 2, 4, 1, 1, 2, 3, 4}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC081Cその2(tt.N, tt.K, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
