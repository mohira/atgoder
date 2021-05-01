package p08_ABC049C

import (
	"reflect"
	"testing"
)

// [AGC013A - Sorted Arrays](https://atcoder.jp/contests/agc013/tasks/agc013_a)
func AnswerAGC013Aその2(A []int) int {
	var count int

	for i := 0; i < len(A); i++ {
		// 取り出せる数列のうち、最も長い数列
		subArray := findGreedyLongSubArray(A[i:])

		// 分割数列分だけindexをずらす(スライスを使うので、-1で補正)
		i += len(subArray) - 1
		count++
	}

	return count
}

func findGreedyLongSubArray(A []int) []int {
	subArrayDesc := getSubArrayDesc(A)
	subArrayAsc := getSubArrayAsc(A)

	if len(subArrayAsc) <= len(subArrayDesc) {
		return subArrayDesc
	} else {
		return subArrayAsc
	}
}

func getSubArrayAsc(A []int) []int {
	// 空のスライスはこない前提
	var now = A[0]

	for i := 0; i < len(A); i++ {
		a := A[i]
		if now <= a {
			now = a
		} else {
			return A[:i]
		}
	}

	return A
}

func getSubArrayDesc(A []int) []int {
	// 空のスライスはこない前提
	var now = A[0]

	for i := 0; i < len(A); i++ {
		a := A[i]
		if a <= now {
			now = a
		} else {
			return A[:i]
		}
	}

	return A
}

func TestAnswerAGC013Aその2(t *testing.T) {
	tests := []struct {
		name string
		N    int
		A    []int
		want int
	}{
		{"入力例1", 6, []int{1, 2, 3, 2, 2, 1}, 2},
		{"入力例2", 9, []int{1, 2, 1, 2, 1, 2, 1, 2, 1}, 5},
		{"入力例3", 7, []int{1, 2, 3, 2, 1, 999999999, 1000000000}, 3},

		{"減少から始まるパターン", 5, []int{3, 2, 1, 2, 3}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerAGC013Aその2(tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}

func TestGetSubArrayDesc(t *testing.T) {
	tests := []struct {
		name string
		A    []int
		want []int
	}{
		{"→ ➘", []int{2, 2, 1}, []int{2, 2, 1}},
		{"→ ➚", []int{2, 2, 3}, []int{2, 2}},
		{"➚ ➚", []int{1, 2, 3}, []int{1}},
		{"→ → →", []int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got = getSubArrayDesc(tt.A)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}

func TestGetSubArrayAsc(t *testing.T) {
	tests := []struct {
		name string
		A    []int
		want []int
	}{
		{"➚ ➚", []int{1, 2, 3}, []int{1, 2, 3}},
		{"→ ➘", []int{2, 2, 1}, []int{2, 2}},
		{"→ ➚", []int{2, 2, 3}, []int{2, 2, 3}},
		{"→ → →", []int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got = getSubArrayAsc(tt.A)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
