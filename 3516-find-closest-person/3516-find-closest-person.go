func findClosest(x int, y int, z int) int {
	xTime, yTime := math.Abs(float64(x-z)), math.Abs(float64(y-z))
	if xTime > yTime {
		return 2
	} else if xTime < yTime {
		return 1
	} else {
		return 0
	}
}