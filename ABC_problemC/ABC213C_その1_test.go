package ABC_problemC

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func AnswerABC213Cその1(H, W, N int, AB [][]int) string {
	bucketA := make(map[int]int)
	bucketB := make(map[int]int)
	for i := 0; i < N; i++ {
		a := AB[i][0]
		bucketA[a] = 1

		b := AB[i][1]
		bucketB[b] = 1
	}

	uniqueA := make([]int, 0, N)
	uniqueB := make([]int, 0, N)
	for a := range bucketA {
		uniqueA = append(uniqueA, a)
	}
	for b := range bucketB {
		uniqueB = append(uniqueB, b)
	}
	sort.Ints(uniqueA)
	sort.Ints(uniqueB)

	mapA := make(map[int]int)
	mapB := make(map[int]int)

	for i, a := range uniqueA {
		mapA[a] = i + 1
	}
	for i, b := range uniqueB {
		mapB[b] = i + 1
	}

	out := make([]string, N)
	for i := 0; i < N; i++ {
		a, b := AB[i][0], AB[i][1]

		out[i] = fmt.Sprintf("%d %d", mapA[a], mapB[b])
	}

	return strings.Join(out, "\n")
}

func TestAnswerABC213C(t *testing.T) {
	tests := []struct {
		name    string
		H, W, N int
		AB      [][]int
		want    string
	}{
		{
			"重複あり",
			4, 5, 3,
			[][]int{
				{3, 2},
				{2, 5},
				{3, 5},
			},
			"2 1\n1 2\n2 2",
		},
		{
			"入力例1",
			4, 5, 2,
			[][]int{
				{3, 2},
				{2, 5},
			},
			"2 1\n1 2",
		},
		{
			"入力例2",
			1000000000, 1000000000, 10,
			[][]int{
				{1, 1},
				{10, 10},
				{100, 100},
				{1000, 1000},
				{10000, 10000},
				{100000, 100000},
				{1000000, 1000000},
				{10000000, 10000000},
				{100000000, 100000000},
				{1000000000, 1000000000},
			},
			"1 1\n2 2\n3 3\n4 4\n5 5\n6 6\n7 7\n8 8\n9 9\n10 10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC213Cその1(tt.H, tt.W, tt.N, tt.AB)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
