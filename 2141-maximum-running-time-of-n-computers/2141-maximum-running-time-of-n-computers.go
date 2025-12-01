func maxRunTime(n int, batteries []int) int64 {
	full := int64(0)
	for _, b := range batteries {
		full += int64(b)
	}

	runTime := func(time int64) bool {
		batteryTime := int64(0)

		for _, b := range batteries {
			batteryTime += min(int64(b), time)
		}

		return batteryTime >= time*int64(n)
	}

	lo := int64(1)
	hi := full / int64(n)
	for lo < hi {
		mid := (hi + lo + 1) / 2

		if ok := runTime(mid); ok {
			lo = mid
		} else {
			hi = mid - 1
		}
	}

	return lo
}