const eps = 1e-6

func judgePoint24(cards []int) bool {
	floatCards := []float64{
		float64(cards[0]),
		float64(cards[1]),
		float64(cards[2]),
		float64(cards[3]),
	}
	return operCards(floatCards)
}

func operCards(cards []float64) bool {
	if len(cards) == 1 {
		return math.Abs(cards[0]-24) < eps
	}
	for i := 0; i < len(cards); i++ {
		for j := i + 1; j < len(cards); j++ {
			calValues := []float64{
				cards[i] + cards[j],
				cards[i] - cards[j],
				cards[j] - cards[i],
				cards[i] * cards[j],
			}
			if cards[j] != 0 {
				calValues = append(calValues, float64(cards[i])/float64(cards[j]))
			}
			if cards[i] != 0 {
				calValues = append(calValues, float64(cards[j])/float64(cards[i]))
			}

			var tempCards []float64
			for z := 0; z < len(cards); z++ {
				if z != i && z != j {
					tempCards = append(tempCards, cards[z])
				}
			}

			for _, calValue := range calValues {
				if operCards(append(tempCards, calValue)) {
					return true
				}
			}
		}
	}
	return false
}