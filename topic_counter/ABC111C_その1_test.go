package topic_counter

import (
	"sort"
	"testing"
)

// [ABC111C \- /\\/\\/\\/](https://atcoder.jp/contests/abc111/tasks/arc103_a)
func AnswerABC111Cその1(N int, V []int) int {
	// 整数が1種類しかない場合は 半分(=N/2)の書き換えが必要
	if isUnique(V) {
		return N / 2
	}

	// 整数が2種類以上ある場合
	odds, evens := divideNumbers(N, V)

	// それぞれのバケットを作成する
	oddsBucket := createBucket(odds)
	evensBucket := createBucket(evens)

	// 最頻値を取得
	oddMode := findMode(oddsBucket)
	evenMode := findMode(evensBucket)

	// 最頻値が異なる場合
	// それぞれの最頻値を残せばOK(⇔最頻値とは異なる整数の頻度を調べればよい)
	if oddMode != evenMode {
		oddOperationCount := len(odds) - findHighestFreq(oddsBucket, 1)
		evenOperationCount := len(evens) - findHighestFreq(evensBucket, 1)

		return oddOperationCount + evenOperationCount
	}

	// 最頻値が一致する場合
	// 奇数列か偶数列どちらかが最頻値をあきらめる
	// → 2番めに大きい頻度を使う(頻度情報だけでいい。どの数字かは不要)
	if oddMode == evenMode {
		opCount1 := count奇数列の最頻値をあきらめた場合の書き換え回数(odds, oddsBucket, evens, evensBucket)
		opCount2 := count偶数列の最頻値をあきらめた場合の書き換え回数(odds, oddsBucket, evens, evensBucket)

		return min(opCount1, opCount2)
	}

	return -99999999
}

func isUnique(V []int) bool {
	v := V[0]
	for i := 1; i < len(V); i++ {
		if v != V[i] {
			return false
		}
	}
	return true
}

func divideNumbers(N int, V []int) ([]int, []int) {
	var odds = make([]int, 0, N/2)
	var evens = make([]int, 0, N/2)

	for i := 0; i < N; i++ {
		v := V[i]
		if i%2 == 0 {
			odds = append(odds, v)
		} else {
			evens = append(evens, v)
		}
	}

	return odds, evens
}

func createBucket(nums []int) map[int]int {
	var bucket = make(map[int]int)
	for _, num := range nums {
		bucket[num]++
	}
	return bucket
}

func findMode(oddsBucket map[int]int) int {
	var mode int
	var maxFreq int
	for num, freq := range oddsBucket {
		if maxFreq < freq {
			maxFreq = freq
			mode = num
		}
	}
	return mode
}

func findHighestFreq(bucket map[int]int, rank int) int {
	var nums = make([]int, len(bucket))

	for _, freq := range bucket {
		nums = append(nums, freq)
	}

	sort.Ints(nums)

	return nums[len(nums)-rank]
}

func count偶数列の最頻値をあきらめた場合の書き換え回数(odds []int, oddsBucket map[int]int, evens []int, evensBucket map[int]int) int {
	oddOperationCount := len(odds) - findHighestFreq(oddsBucket, 1)
	evenOperationCount := len(evens) - findHighestFreq(evensBucket, 2)

	return oddOperationCount + evenOperationCount
}

func count奇数列の最頻値をあきらめた場合の書き換え回数(odds []int, oddsBucket map[int]int, evens []int, evensBucket map[int]int) int {
	oddOperationCount := len(odds) - findHighestFreq(oddsBucket, 2)
	evenOperationCount := len(evens) - findHighestFreq(evensBucket, 1)

	return oddOperationCount + evenOperationCount
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func TestAnswerABC111Cその1(t *testing.T) {
	tests := []struct {
		name string
		N    int
		V    []int
		want int
	}{
		{"入力例0: 奇数番目も偶数番目も書き換えが必要", 6, []int{3, 1, 3, 2, 8, 9}, 3},
		{"入力例1: 奇数番目は条件を満たしている", 4, []int{3, 1, 3, 2}, 1},
		{"入力例2: 最初から条件を満たしているので書き換え不要な場合", 6, []int{105, 119, 105, 119, 105, 119}, 0},
		{"入力例3: 1種類しかない場合", 4, []int{1, 1, 1, 1}, 2},

		{"奇数列と偶数列の最頻値が一致する場合", 6, []int{1, 1, 2, 2, 2, 2}, 3},
		{"奇数列と偶数列の最頻値が一致する場合(どちらの数列も 3 が最頻)で奇数列を優遇", 10, []int{3, 1, 3, 2, 3, 3, 1, 3, 1, 4}, 6},
		{"奇数列と偶数列の最頻値が一致する場合(どちらの数列も 3 が最頻)で偶数列を優遇", 10, []int{1, 3, 2, 3, 3, 3, 3, 1, 4, 1}, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC111Cその1(tt.N, tt.V)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
