func getHappyString(n int, k int) string {
	total := 3 * (1 << (n - 1))
	if k > total {
		return ""
	}

	res := make([]byte, 0, n)
	prev := byte(0)

	for i := 0; i < n; i++ {
		for ch := byte('a'); ch <= 'c'; ch++ {
			if ch == prev {
				continue
			}

			remain := n - i - 1
			cnt := 1 << remain

			if k > cnt {
				k -= cnt
			} else {
				res = append(res, ch)
				prev = ch
				break
			}
		}
	}
	return string(res)
}