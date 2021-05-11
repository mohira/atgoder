package topic_binary_search

import (
	"fmt"
	"testing"
)

// p.094 code6.1 配列から目的の値を探索する二分探索法
// 特定の値のindexを返す
// 見つからない場合は -1 を返す
func code61(target int) int {
	var A = []int{3, 5, 8, 10, 14, 17, 21, 39}
	left := 0
	right := len(A) - 1

	for left <= right {
		mid := (left + right) / 2 // 区間の真ん中
		fmt.Printf("left=%d mid=%d right=%d v=%d\n", left, mid, right, A[mid])
		fmt.Println(A[left : right+1])

		if A[mid] == target {
			return mid
		} else if target < A[mid] {
			right = mid - 1 // 探索範囲を左に絞る
		} else if A[mid] < target {
			left = mid + 1 // 探索範囲を右に絞る
		}
	}

	return -1
}

func TestKenchonBookBinarySearchCode61(t *testing.T) {
	tests := []struct {
		name   string
		target int
		want   int
	}{
		{"ex1: 1回で求まる", 10, 3},
		{"ex2: 真ん中より小さい", 3, 0},
		{"ex3: 真ん中より大きい", 39, 7},
		{"存在しないケース", -100, -1},
		{"存在しないケース", 9, -1},
		{"存在しないケース", 100, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := code61(tt.target)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
