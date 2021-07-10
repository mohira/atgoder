package topic_dp

import "math"

func AbsInt(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func Min(nums ...int) int {
	min := math.MaxInt64

	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	return min
}
