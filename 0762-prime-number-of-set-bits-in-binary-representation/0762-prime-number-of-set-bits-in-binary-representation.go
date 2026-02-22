func countPrimeSetBits(left int, right int) int {
	primeSet := map[int]struct{}{2: {}, 3: {}, 5: {}, 7: {}, 11: {}, 13: {}, 17: {}, 19: {}, 23: {}, 29: {}, 31: {}}

	countOneBit := func(n int) int {
		cnt := 0
		for n > 0 {
			if n&1 == 1 {
				cnt++
			}
			n >>= 1
		}
		return cnt
	}

	res := 0
	for i := left; i <= right; i++ {
		oneCnt := countOneBit(i)
		if _, ok := primeSet[oneCnt]; ok {
			res++
		}
	}

	return res
}