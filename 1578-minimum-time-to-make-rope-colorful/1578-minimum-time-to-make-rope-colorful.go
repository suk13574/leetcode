func minCost(colors string, neededTime []int) int {
	res := 0

	var nowColor string
	nowNeedTime := 0
	nowMaxTime := 0
	for i := 0; i < len(colors); i++ {
		color := string(colors[i])
		if nowColor == color {
			nowNeedTime += neededTime[i]
			nowMaxTime = max(nowMaxTime, neededTime[i])
		} else {
			res += nowNeedTime - nowMaxTime

			nowColor = color
			nowNeedTime = neededTime[i]
			nowMaxTime = neededTime[i]
		}
	}

	return res + nowNeedTime - nowMaxTime
}