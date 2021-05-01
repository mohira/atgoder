package p08_ABC049C

import (
	"reflect"
	"testing"
)

// [AGC013A - Sorted Arrays](https://atcoder.jp/contests/agc013/tasks/agc013_a)
func AnswerAGC013Aその1(N int, A []int) int {
	var countSubArrays int

	for {
		subArrayDesc, subArrayAsc := subArray単調非減少(A), subArray単調非増加(A)
		lengthDesc, lengthAsc := len(subArrayDesc), len(subArrayAsc)

		countSubArrays++

		// 分割した数列の長さ が もとの数例Aの長さ と一致 => 終了
		lengthA := len(A)
		if (lengthA == lengthDesc) || (lengthA == lengthAsc) {
			break
		}

		// 分割した数列を切り出す
		// 長さが同じ場合はどっちでやってもいい
		if lengthDesc <= lengthAsc {
			A = A[lengthAsc:]
		} else {
			A = A[lengthDesc:]
		}
	}

	return countSubArrays
}

func getDiffs(nums []int) []int {
	var diffs = make([]int, 0, len(nums)-1)

	for i := 0; i < len(nums)-1; i++ {
		d := nums[i+1] - nums[i]
		diffs = append(diffs, d)
	}

	return diffs
}

func subArray単調非増加(nums []int) []int {
	var i int

	for _, n := range getDiffs(nums) {
		if 0 <= n {
			i++
		} else {
			break
		}
	}

	return nums[:i+1]
}

func subArray単調非減少(nums []int) []int {
	var i int
	for _, n := range getDiffs(nums) {
		if n <= 0 {
			i++
		} else {
			break
		}

	}
	return nums[:i+1]
}

func TestAnswerAGC013A(t *testing.T) {
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
			got := AnswerAGC013Aその1(tt.N, tt.A)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}

func TestSubArray単調非減少(t *testing.T) {
	tests := []struct {
		name string
		A    []int
		want []int
	}{
		{"→ ➘", []int{2, 2, 1}, []int{2, 2, 1}},
		{"→ ➚", []int{2, 2, 3}, []int{2, 2}},
		{"➚ ➚", []int{1, 2, 3}, []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got = subArray単調非減少(tt.A)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}

func TestSubArray単調非増加(t *testing.T) {
	tests := []struct {
		name string
		A    []int
		want []int
	}{
		{"➚ ➚", []int{1, 2, 3}, []int{1, 2, 3}},
		{"→ ➘", []int{2, 2, 1}, []int{2, 2}},
		{"→ ➚", []int{2, 2, 3}, []int{2, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got = subArray単調非増加(tt.A)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
