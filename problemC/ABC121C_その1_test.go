package problemC

import (
	"atgoder/lib"
	"sort"
	"testing"
)

type DrinkShop struct {
	UnitPrice int
	Stock     int
}

type DrinkShops []DrinkShop

func (ds DrinkShops) Len() int {
	return len(ds)
}

func (ds DrinkShops) Less(i, j int) bool {
	if ds[i].UnitPrice < ds[j].UnitPrice {
		return true
	}
	return false
}

func (ds DrinkShops) Swap(i, j int) {
	ds[i], ds[j] = ds[j], ds[i]
}

// [ABC121C - Energy Drink Collector](https://atcoder.jp/contests/abc121/tasks/abc121_c)
func AnswerABC121Cその1(N, M int, AB [][]int) int {
	// 栄養ドリンクが安いところから貪欲に買っていく
	shops := make(DrinkShops, 0, N)
	for i := 0; i < N; i++ {
		a, b := AB[i][0], AB[i][1]
		shops = append(shops, DrinkShop{a, b})
	}
	sort.Stable(shops)

	totalQuantity := 0
	cost := 0

	for _, shop := range shops {
		quantity := lib.Min(M-totalQuantity, shop.Stock)

		cost += shop.UnitPrice * quantity
		totalQuantity += quantity

		if totalQuantity == M {
			break
		}
	}

	return cost
}

func TestAnswerABC121Cその1(t *testing.T) {
	tests := []struct {
		name string
		N, M int
		AB   [][]int
		want int
	}{
		{
			"入力例1",
			2, 5,
			[][]int{
				{4, 9},
				{2, 4},
			},
			12,
		},
		{
			"入力例2",
			4, 30,
			[][]int{
				{6, 18},
				{2, 5},
				{3, 10},
				{7, 9},
			},
			130,
		},
		{
			"入力例3",
			1, 100000,
			[][]int{
				{1000000000, 100000},
			},
			100000000000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC121Cその1(tt.N, tt.M, tt.AB)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
