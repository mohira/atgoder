package problemB

import (
	"sort"
	"strconv"
	"testing"
)

type Review struct {
	RestaurantId int
	Name         string
	Point        int
}

type Reviews []Review

func (rs Reviews) Len() int {
	return len(rs)
}

// 辞書順でソート。名前が等しいなら得点降順。
func (rs Reviews) Less(i, j int) bool {
	if rs[i].Name < rs[j].Name {
		return true
	}
	if (rs[i].Name == rs[j].Name) && (rs[i].Point > rs[j].Point) {
		return true
	}

	return false
}

func (rs Reviews) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

// [ABC128B - Guidebook](https://atcoder.jp/contests/abc130/tasks/abc130_b)
func AnswerABC128Bその1(N int, S []string, P []int) string {
	reviews := make(Reviews, N)

	for i := 0; i < N; i++ {
		reviews[i] = Review{
			RestaurantId: i + 1,
			Name:         S[i],
			Point:        P[i],
		}
	}

	sort.Stable(reviews)

	var ans string
	for _, review := range reviews {
		id := strconv.Itoa(review.RestaurantId)
		if ans == "" {
			ans += id
		} else {
			ans += "\n" + id
		}
	}
	return ans
}

func TestAnswerABC128Bその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		S    []string
		P    []int
		want string
	}{
		{
			"入力例1",
			6,
			[]string{"khabarovsk", "moscow", "kazan", "kazan", "moscow", "khabarovsk"},
			[]int{20, 10, 50, 35, 60, 40},
			"3\n4\n6\n1\n5\n2",
		},
		{
			"入力例2",
			10,
			[]string{"yakutsk", "yakutsk", "yakutsk", "yakutsk", "yakutsk", "yakutsk", "yakutsk", "yakutsk", "yakutsk", "yakutsk"},
			[]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
			"10\n9\n8\n7\n6\n5\n4\n3\n2\n1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC128Bその1(tt.N, tt.S, tt.P)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
