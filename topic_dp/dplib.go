package topic_dp

func AbsInt(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func ChMin(a *int, b int) {
	if b < *a {
		*a = b
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
