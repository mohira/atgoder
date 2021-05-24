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

func sumInts(nums []int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	return total
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

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(integers []int) int {
	a := integers[0]
	b := integers[1]
	result := a * b / GCD(a, b)

	for _, integer := range integers[2:] {
		result = LCM([]int{result, integer})
	}

	return result
}

// ある整数の桁数を返す
func getDigitNums(N int) int {
	var digit int
	for N > 0 {
		N /= 10
		digit++
	}
	return digit
}
