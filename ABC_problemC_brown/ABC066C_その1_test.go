package ABC_problemC_brown

import (
	"container/list"
	"fmt"
	"strings"
	"testing"
)

// [ABC066C - pushpush](https://atcoder.jp/contests/abc066/tasks/abc066_c)
func AnswerABC066Cその1(N int, A []int) string {
	// 反転操作を伴う場合はデキューで対応する
	l := list.New()

	for i := 0; i < N; i++ {
		a := A[i]
		if i%2 == 1 {
			l.PushBack(a)
		} else {
			l.PushFront(a)
		}
	}

	var s = make([]string, 0, N)

	// 反転回数の偶奇で最終出力が変わる
	// 偶数回反転 → 逆順出力
	// 奇数回反転 → 　順出力
	if N%2 == 0 {
		for e := l.Back(); e != nil; e = e.Prev() {
			s = append(s, fmt.Sprintf("%d", e.Value))
		}
	} else {
		for e := l.Front(); e != nil; e = e.Next() {
			s = append(s, fmt.Sprintf("%d", e.Value))
		}
	}

	return strings.Join(s, " ")
}

func TestAnswerABC066Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want string
	}{
		{"入力例1", 4, []int{1, 2, 3, 4}, "4 2 1 3"},
		{"入力例2", 3, []int{1, 2, 3}, "3 1 2"},
		{"入力例3", 1, []int{1000000000}, "1000000000"},
		{"入力例4", 6, []int{0, 6, 7, 6, 7, 0}, "0 6 6 0 7 7"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC066Cその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
