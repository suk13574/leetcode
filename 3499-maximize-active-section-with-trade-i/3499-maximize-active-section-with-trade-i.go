func maxActiveSectionsAfterTrade(s string) int {
	totalOne := 0
	prevZero := -1
	currentZero := 0
	maxGain := 0

	for i := 0; i <= len(s); i++ {
		var b byte

		if i == len(s) {
			b = '1'
		} else {
			b = s[i]
		}

		if b == '0' {
			currentZero++
			continue
		}

        if i < len(s) {
		    totalOne++
        }

		if currentZero > 0 {
            if prevZero != -1 {
                maxGain = max(maxGain, currentZero+prevZero)
            }
            prevZero = currentZero
		    currentZero = 0
		}
	}

	return maxGain + totalOne
}