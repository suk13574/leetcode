func sumFourDivisors(nums []int) int {
	res := 0
	for _, n := range nums {
		for p := 2; p*p <= n; p++ {
			if n%p != 0 {
				continue
			}

			q := n / p

			if q != p && isPrime(p) && isPrime(q) {
				res += 1 + p + q + n
				break
			}

			if n == p*p*p && isPrime(p) {
				res += 1 + p + p*p + n
			}
		}

	}
	return res
}

func isPrime(x int) bool {
	if x < 2 {
		return false
	}

	if x%2 == 0 {
		return x == 2
	}

	for i := 3; i*i <= x; i += 2 {
		if x%i == 0 {
			return false
		}
	}

	return true
}