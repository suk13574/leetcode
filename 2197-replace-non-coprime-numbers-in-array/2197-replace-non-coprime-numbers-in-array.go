func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func replaceNonCoprimes(nums []int) []int {
	top := -1
	for _, b := range nums {
		for top >= 0 {
			a := nums[top]
			g := gcd(a, b)
			if g == 1 {
				break
			}

			lcm := int((int64(a) / int64(g)) * int64(b))
			b = lcm
			top--
		}
		top++
		nums[top] = b
	}
	return nums[:top+1]
}