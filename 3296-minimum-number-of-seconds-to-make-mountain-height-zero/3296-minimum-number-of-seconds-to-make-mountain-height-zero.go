func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	sort.Slice(workerTimes, func(i, j int) bool { return workerTimes[i] <= workerTimes[j] })

	can := func(t int64) bool {
		total := int64(0)

		for _, wt := range workerTimes {
			wwt := int64(wt)

			// find x
			lo, hi := int64(0), int64(mountainHeight)
			for lo < hi {
				mid := lo + (hi-lo+1)/2
				xt := wwt * mid * (mid + 1) / 2

				if xt <= t {
					lo = mid
				} else {
					hi = mid - 1
				}
			}

			total += lo

			if total >= int64(mountainHeight) {
				return true
			}
		}
		return false
	}

	lo := int64(1)
	hi := (int64(workerTimes[0]) * int64(mountainHeight) * int64(mountainHeight+1)) / 2

	for lo < hi {
		mid := lo + (hi-lo)/2
		if can(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return lo
}
