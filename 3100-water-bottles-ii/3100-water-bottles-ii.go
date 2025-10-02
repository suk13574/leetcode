func maxBottlesDrunk(numBottles int, numExchange int) int {
	res := numBottles

	for numBottles >= numExchange {
		numBottles = numBottles - numExchange + 1
		res++
		numExchange++
	}

	return res
}