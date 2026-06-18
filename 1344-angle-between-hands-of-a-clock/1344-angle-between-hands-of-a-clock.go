func angleClock(hour int, minutes int) float64 {
	minutesAngle := float64(minutes * 6)
	hourAngle := float64(hour*30) + float64(minutes)*0.5

	diff := abs(minutesAngle - hourAngle)

	return min(diff, 360-diff)
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}