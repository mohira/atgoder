package topic_cumlative_sum

import (
	"sort"
	"testing"
)

type Event struct {
	when int
	cost int
}

// [ABC188D - Snuke Prime](https://atcoder.jp/contests/abc188/tasks/abc188_d)
func AnswerABC188Dその1(N, C int, ABC [][]int) int {
	events := make([]Event, 0, N)

	for i := 0; i < N; i++ {
		a, b, c := ABC[i][0], ABC[i][1], ABC[i][2]
		a--
		events = append(events, Event{a, +c})
		events = append(events, Event{b, -c})
	}
	// https://yourbasic.org/golang/how-to-sort-in-go/
	sort.SliceStable(events, func(i, j int) bool {
		return events[i].when < events[j].when
	})

	var ans int
	var fee int
	var t int
	for _, event := range events {
		if event.when != t {
			ans += min(C, fee) * (event.when - t)
			t = event.when
		}
		fee += event.cost
	}

	return ans
}

func TestAnswerABC188Dその1(t *testing.T) {
	tests := []struct {
		name string
		N, C int
		ABC  [][]int
		want int
	}{
		{"入力例1",
			2, 6,
			[][]int{
				{1, 2, 4},
				{2, 2, 4},
			},
			10},

		{"入力例2",
			5, 1000000000,
			[][]int{
				{583563238, 820642330, 44577},
				{136809000, 653199778, 90962},
				{54601291, 785892285, 50554},
				{5797762, 453599267, 65697},
				{468677897, 916692569, 87409},
			},
			163089627821228},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC188Dその1(tt.N, tt.C, tt.ABC)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})

	}

}
