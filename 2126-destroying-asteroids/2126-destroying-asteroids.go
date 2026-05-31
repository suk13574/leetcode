func asteroidsDestroyed(mass int, asteroids []int) bool {
	sort.Ints(asteroids)

	total := int64(mass)
	for _, a := range asteroids {
		if total < int64(a) {
			return false
		}

		total += int64(a)
	}

	return true
}