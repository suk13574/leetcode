func minimumDistance(nums []int) int {
	set := make(map[int][]int)

	for i, num := range nums {
		if _, ok := set[num]; !ok {
			set[num] = []int{}
		}
		set[num] = append(set[num], i)
	}

	res := 1 << 31
	for _, idxs := range set {
		if len(idxs) < 3 {
			continue
		}
		for i := 0; i+2 < len(idxs); i++ {
			d := (idxs[i+2] - idxs[i]) * 2
			res = min(res, d)
		}
	}

	if res == 1<<31 {
		return -1
	}
	return res
}