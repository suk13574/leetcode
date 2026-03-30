func checkStrings(s1 string, s2 string) bool {
	n := len(s1)
	even := make(map[byte]int)
	odd := make(map[byte]int)

	for i := 0; i < n; i++ {
		if i&1 == 1 {
			odd[s1[i]]++
		} else {
			even[s1[i]]++
		}
	}

	for i := 0; i < n; i++ {
		if i&1 == 1 {
			if _, ok := odd[s2[i]]; !ok {
				return false
			}
			odd[s2[i]]--
			if odd[s2[i]] == 0 {
				delete(odd, s2[i])
			}
		} else {
			if _, ok := even[s2[i]]; !ok {
				return false
			}
			even[s2[i]]--
			if even[s2[i]] == 0 {
				delete(even, s2[i])
			}
		}
	}

	if len(odd) > 0 || len(even) > 0 {
		return false
	}

	return true
}