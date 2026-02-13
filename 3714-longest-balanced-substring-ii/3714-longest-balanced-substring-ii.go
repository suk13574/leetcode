func longestBalancedOneChar(s string) int {
	if len(s) == 0 {
		return 0
	}

	res := 1
	cur := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cur++
		} else {
			cur = 1
		}
		res = max(res, cur)
	}
	return res
}

func longestBalancedPair(s string, ch1, ch2 byte) int {
	res := 0

	start := 0
	first := map[int]int{0: start - 1}

	diff := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ch1 {
			diff--
		} else if s[i] == ch2 {
			diff++
		} else {
			start = i + 1
			diff = 0
			first = map[int]int{0: start - 1}
			continue
		}

		if idx, ok := first[diff]; ok {
			res = max(res, i-idx)
		} else {
			first[diff] = i
		}
	}

	return res
}

func longestBalancedABC(s string) int {
	n := len(s)

	offset := n
	base := 2*n + 1

	encode := func(d1, d2 int) int64 {
		return int64(d1+offset)*int64(base) + int64(d2+offset)
	}

	a, b, c := 0, 0, 0
	first := make(map[int64]int, n+1)
	first[encode(0, 0)] = -1

	res := 0
	for i := 0; i < n; i++ {
		switch s[i] {
		case 'a':
			a++
		case 'b':
			b++
		case 'c':
			c++
		}

		key := encode(a-b, a-c)
		if idx, ok := first[key]; ok {
			res = max(res, i-idx)
		} else {
			first[key] = i
		}
	}

	return res
}

func longestBalanced(s string) int {
	res := 0
	res = max(res, longestBalancedOneChar(s))

	res = max(res, longestBalancedPair(s, 'a', 'b'))
	res = max(res, longestBalancedPair(s, 'a', 'c'))
	res = max(res, longestBalancedPair(s, 'b', 'c'))

	res = max(res, longestBalancedABC(s))

	return res
}