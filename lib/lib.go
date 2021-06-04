package lib

import (
	"errors"
	"math"
)

func AbsInt(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func Min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func Max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func FindMin(nums []int) int {
	var min = math.MaxInt32
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return min
}

func FindMax(nums []int) int {
	var max = -math.MaxInt32
	for _, n := range nums {
		if max < n {
			max = n
		}
	}
	return max
}

// SumIntRange 整数区間の和を求める
//   first last
//   ↓     ↓
//   3 4 5 6 => 18
func SumIntRange(first, last int) int {
	return SumFrom1toN(last) - SumFrom1toN(first-1)
}

func SumFrom1toN(n int) int {
	return n * (n + 1) / 2
}

func SumInts(nums []int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}

func IsPalindrome(s string) bool {
	return ReverseStr(s) == s
}

func ReverseStr(s string) string {
	var rs string

	for i := len(s) - 1; i >= 0; i-- {
		rs += string(s[i])
	}

	return rs
}

func IsLower(r rune) bool {
	return 'a' <= r && r <= 'z'
}

func IsUpper(r rune) bool {
	return 'A' <= r && r <= 'Z'
}

// GCD 2つの整数の最大公約数を求める。ユークリッドの互除法で。
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM 整数スライスの最小公倍数を求める
func LCM(integers []int) int {
	a := integers[0]
	b := integers[1]
	result := a * b / GCD(a, b)

	for _, integer := range integers[2:] {
		result = LCM([]int{result, integer})
	}

	return result
}

// GetDigitNums ある整数の桁数を返す
func GetDigitNums(N int) int {
	var digit int
	for N > 0 {
		N /= 10
		digit++
	}
	return digit
}

// GetNumberOfDigits 非負の整数の桁数を求める
func GetNumberOfDigits(n int) (int, error) {
	switch {
	case n < 0:
		return 0, errors.New("マイナスは認めてないよ")
	case n == 0:
		return 1, nil
	}

	digit := 0
	for n > 0 {
		n /= 10
		digit++
	}
	return digit, nil
}
