func distance(nums []int) []int64 {
	idxsOfNum := make(map[int][]int)
	for i, v := range nums {
		idxsOfNum[v] = append(idxsOfNum[v], i)
	}

	res := make([]int64, len(nums))
	for _, idxs := range idxsOfNum {
		k := len(idxs)
		prefix := make([]int64, k+1)
		for p, idx := range idxs {
			prefix[p+1] = prefix[p] + int64(idx)
		}

		for p, idx := range idxs {
			leftSum := int64(p)*int64(idx) - prefix[p]
			rightSum := (prefix[k] - prefix[p+1]) - int64(k-p-1)*int64(idx)
			res[idx] = leftSum + rightSum
		}
	}

	return res
}