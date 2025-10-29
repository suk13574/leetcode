func smallestNumber(n int) int {
	res := 0
	nn := n
	for nn > 0 {
		nn = nn >> 1
		res = res*2 + 1
	}
	return res
}