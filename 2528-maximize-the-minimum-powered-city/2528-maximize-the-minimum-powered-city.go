func maxPower(stations []int, r int, k int) int64 {
	n := len(stations)
	initPower := make([]int64, n)

	var sum int64
	for i := 0; i < r && i < n; i++ {
		sum += int64(stations[i])
	}

	for i := 0; i < n; i++ {
		if i-r-1 >= 0 {
			sum -= int64(stations[i-r-1])
		}
		if i+r < n {
			sum += int64(stations[i+r])
		}
		initPower[i] = sum
	}

	minInit := initPower[0]
	for i := 1; i < n; i++ {
		if initPower[i] < minInit {
			minInit = initPower[i]
		}
	}
	lo := int64(0)
	hi := minInit + int64(k)

	// check if all cities can reach >= x
	// sweep left to right
	feasible := func(x int64) bool {
		leftK := int64(k)
		diff := make([]int64, n+1)
		var runningExtra int64

		for i := 0; i < n; i++ {
			runningExtra += diff[i]
			cur := initPower[i] + runningExtra
			if cur < x {
				need := x - cur
				leftK -= need
				if leftK < 0 {
					return false
				}

				runningExtra += need
				end := i + 2*r
				if end+1 < n { // expire added power after end
					diff[end+1] -= need
				}
			}
		}
		return true
	}

	// binary search
	for lo < hi {
		mid := (lo + hi + 1) / 2 // upper mid
		if feasible(mid) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return lo
}