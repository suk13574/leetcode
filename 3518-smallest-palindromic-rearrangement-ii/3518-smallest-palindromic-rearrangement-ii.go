func smallestPalindrome(s string, k int) string {
	n := len(s)
	half := n / 2

	count := make([]int, 26)
	for i := 0; i < half; i++ {
		count[int(s[i]-'a')]++
	}

	if countPermutations(count, int64(k)) < int64(k) {
		return ""
	}

	left := make([]byte, half)

	for pos := 0; pos < half; pos++ {
		picked := false

		for ch := 0; ch < 26; ch++ {
			if count[ch] == 0 {
				continue
			}

			count[ch]--

			ways := countPermutations(count, int64(k))

			if ways >= int64(k) {
				left[pos] = byte('a' + ch)
				picked = true
				break
			}

			k -= int(ways)
			count[ch]++
		}

		if !picked {
			return ""
		}
	}

	res := make([]byte, n)
	copy(res, left)

	if n&1 == 1 {
		res[half] = s[half]
	}

	for i := 0; i < half; i++ {
		res[n-1-i] = res[i]
	}

	return string(res)
}

func countPermutations(count []int, limit int64) int64 {
	ways := int64(1)
	used := 0

	for _, c := range count {
		if c == 0 {
			continue
		}

		combLimit := (limit + ways - 1) / ways
		comb := combinationCap(used+c, c, combLimit)

		if comb >= combLimit {
			return limit
		}

		ways *= comb
		used += c
	}

	return ways
}

func combinationCap(n, r int, limit int64) int64 {
	if r > n-r {
		r = n - r
	}

	res := int64(1)

	for i := 1; i <= r; i++ {
		numerator := int64(n - r + i)
		denomiator := int64(i)

		res = res * numerator / denomiator

		if res >= limit {
			return limit
		}
	}

	return res
}
