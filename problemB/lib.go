package problemB

import "math"

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
