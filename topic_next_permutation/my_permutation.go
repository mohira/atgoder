package topic_next_permutation

// [Go言語でスライスの要素の順列、組み合わせを与える \- Camera Obscura](http://obelisk.hatenablog.com/entry/2018/10/09/025240)
func remove(ar []int, i int) []int {
	tmp := make([]int, len(ar))
	copy(tmp, ar)
	return append(tmp[0:i], tmp[i+1:]...)
}

func permutation(intSlice []int) [][]int {
	var result [][]int
	if len(intSlice) == 1 {
		return append(result, intSlice)
	}
	for i, a := range intSlice {
		for _, b := range permutation(remove(intSlice, i)) {
			result = append(result, append([]int{a}, b...))
		}
	}
	return result
}
