package problemB

import "math"

func absInt(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func findMin(nums []int) int {
	var min = math.MaxInt32
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return min
}

func findMax(nums []int) int {
	var max = -math.MaxInt32
	for _, n := range nums {
		if max < n {
			max = n
		}
	}
	return max
}

// 整数区間の和を求める
//   first last
//   ↓     ↓
//   3 4 5 6 => 18
func sumIntRange(first, last int) int {
	return sumFrom1toN(last) - sumFrom1toN(first-1)
}

func sumFrom1toN(n int) int {
	return n * (n + 1) / 2
}

func isPalindrome(s string) bool {
	return reverseStr(s) == s
}

func reverseStr(s string) string {
	var rs string

	for i := len(s) - 1; i >= 0; i-- {
		rs += string(s[i])
	}

	return rs
}

func isLower(r rune) bool {
	return 'a' <= r && r <= 'z'
}

func isUpper(r rune) bool {
	return 'A' <= r && r <= 'Z'
}
