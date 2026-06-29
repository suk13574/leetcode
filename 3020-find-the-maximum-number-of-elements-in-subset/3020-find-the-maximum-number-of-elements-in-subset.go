func maximumLength(nums []int) int {
	numCnt := make(map[int]int)

	for _, v := range nums {
		numCnt[v]++
	}

	res := 1

	if v, ok := numCnt[1]; ok {
		if v%2 == 0 {
			v--
		}
		res = max(res, v)
	}

	for num := range numCnt {
		if num == 1 {
			continue
		}

		length := 1
		cur := num

		for numCnt[cur] >= 2 {
			if cur > 1_000_000_000/cur {
				break
			}

			next := cur * cur

			if numCnt[next] == 0 {
				break
			}

			length += 2
			cur = next
		}

		res = max(res, length)
	}

	return res
}