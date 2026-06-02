func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	n := len(landDuration)
	m := len(waterDuration)

	res := 1 << 30
	for i := 0; i < n; i++ {
		initT := landStartTime[i] + landDuration[i]
		for j := 0; j < m; j++ {
			start := waterStartTime[j]
			totalT := initT

			if start > totalT {
				totalT += start - initT
			}
			totalT += waterDuration[j]
			res = min(res, totalT)
		}
	}

	for j := 0; j < m; j++ {
		initT := waterStartTime[j] + waterDuration[j]
		for i := 0; i < n; i++ {
			start := landStartTime[i]
			totalT := initT

			if start > totalT {
				totalT += start - initT
			}
			totalT += landDuration[i]
			res = min(res, totalT)
		}
	}

	return res
}