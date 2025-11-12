func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func minOperations(nums []int) int {
	n := len(nums)

	// If the gcd of the entire array is not 1, it's impossible.
	all := 0
	for _, num := range nums {
		all = gcd(all, num)
	}
	if all != 1 {
		return -1
	}

	// If the array already contains 1, answer is n - count(1)
	ones := 0
	for _, num := range nums {
		if num == 1 {
			ones++
		}
	}
	if ones > 0 {
		return n - (ones)
	}

	// Find the shortest subarray whose gcd is 1.
	minLen := n + 1
	for i := 0; i < n; i++ {
		g := nums[i]
		for j := i + 1; j < n; j++ {
			g = gcd(g, nums[j])
			if g == 1 {
				minLen = min(minLen, j-i+1)
				break
			}
		}
	}

	return (minLen - 1) + (n - 1)
}