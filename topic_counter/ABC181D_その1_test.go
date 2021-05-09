package topic_counter

import (
	"strconv"
	"testing"
)

// [ABC181D - Hachi](https://atcoder.jp/contests/abc181/tasks/abc181_d)
func AnswerABC181Dその1(S string) string {
	// 1桁なら8の倍数チェックでOK
	if len(S) == 1 {
		n, err := strconv.Atoi(S)
		if err != nil {
			panic(err)
		}

		if n%8 == 0 {
			return "Yes"
		} else {
			return "No"
		}
	}

	// 2桁の場合は ab と ba の2通りなので素直にチェック
	if len(S) == 2 {
		a := int(S[0] - '0') // 1桁目
		b := int(S[1] - '0') // 2桁目
		if (10*a+b)%8 == 0 || (10*b+a)%8 == 0 {
			return "Yes"
		} else {
			return "No"
		}
	}

	// 3桁以上の場合
	// Sから生成可能な頻度情報
	sBucket := make([]int, 10)
	for _, s := range S {
		sBucket[s-'0']++
	}

	for _, n8x := range get3桁の8の倍数() {
		n8xBucket := createIntBucket(n8x)

		// 頻度情報の減算
		diffBucket(n8xBucket, sBucket)

		// 頻度情報が0ということは、n8xBucketの構成要素を sBucketが持っている
		// つまり、Sに含まれる3つの数字を入れ替えて n8x を作れることを意味する
		if totalFreq(n8xBucket) == 0 {
			return "Yes"
		}
	}

	return "No"
}

func totalFreq(bucket []int) int {
	var total int

	for _, freq := range bucket {
		total += freq
	}

	return total
}

// 頻度情報の減算
// Pythonの Counter.__sub__() のような実装
func diffBucket(bucket1, bucket2 []int) []int {
	N := len(bucket1)

	for i := 0; i < N; i++ {
		if bucket1[i]-bucket2[i] <= 0 {
			bucket1[i] = 0 // 下限は0
		} else {
			bucket1[i] -= bucket2[i]
		}
	}

	return bucket1
}

func createIntBucket(N int) []int {
	bucket := make([]int, 10)

	for _, s := range strconv.Itoa(N) {
		bucket[s-'0']++
	}

	return bucket
}

func get3桁の8の倍数() []int {
	var nums []int

	for i := 104; i < 1000; i += 8 {
		nums = append(nums, i)
	}

	return nums
}

func TestAnswerABC181Dその1(t *testing.T) {
	tests := []struct {
		name string
		S    string
		want string
	}{
		// 3桁以上
		{"3桁でYes", "821", "Yes"},
		{"3桁でNo", "999", "No"},
		{"4桁でYes", "1234", "Yes"},
		{"4桁でNo", "1333", "No"},

		// 2桁
		{"2桁でそのまま8の倍数", "88", "Yes"},
		{"2桁で入れかえればYes", "42", "Yes"},
		{"2桁で入れかえてもNo", "99", "No"},

		// 1桁
		{"1桁で8の倍数である", "8", "Yes"},
		{"1桁で8の倍数でない", "9", "No"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnswerABC181Dその1(tt.S)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
