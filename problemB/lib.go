package problemB

import "math"

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
