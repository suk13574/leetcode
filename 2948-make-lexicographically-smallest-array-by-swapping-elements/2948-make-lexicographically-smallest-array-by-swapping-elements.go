type Pair struct {
	val int
	idx int
}

func lexicographicallySmallestArray(nums []int, limit int) []int {
	n := len(nums)

	arr := make([]Pair, n)

	for i, num := range nums {
		arr[i] = Pair{num, i}
	}

	sort.Slice(arr, func(i, j int) bool { return arr[i].val < arr[j].val })

	res := make([]int, n)

	start := 0

	for start < n {
		end := start

		for end+1 < n && arr[end+1].val-arr[end].val <= limit {
			end++
		}

		indices := make([]int, 0, end-start+1)

		for i := start; i <= end; i++ {
			indices = append(indices, arr[i].idx)
		}

		sort.Ints(indices)

		for i := 0; i < len(indices); i++ {
			res[indices[i]] = arr[start+i].val
		}

		start = end + 1
	}

	return res
}