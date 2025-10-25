func totalMoney(n int) int {
	w := n / 7
	d := n % 7

	res := 28*w + 7*w*(w-1)/2
	res += d*(w+1) + d*(d-1)/2

	return res
}