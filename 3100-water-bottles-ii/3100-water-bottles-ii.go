func maxBottlesDrunk(numBottles int, numExchange int) int {
	res := numBottles

	for numBottles >= numExchange {
		// bottles := numBottles / numExchange
		// numBottles = bottles + (numBottles % numExchange)

		// res += bottles
		numBottles = numBottles - numExchange + 1
		res++
		numExchange++
	}

	return res
}