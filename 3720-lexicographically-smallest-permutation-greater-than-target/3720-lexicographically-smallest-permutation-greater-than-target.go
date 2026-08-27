func lexGreaterPermutation(s string, target string) string {
	n := len(s)

	have := [26]int{}
	for _, ch := range s {
		have[ch-'a']++
	}

	pos := 0
	res := make([]byte, 0, n)

	for pos < n {
		ch := target[pos]

		if have[ch-'a'] == 0 {
			break
		}

		res = append(res, ch)
		have[ch-'a']--
		pos++
	}

	if pos == n {
		pos--
		last := res[len(res)-1]
		res = res[:len(res)-1]
		have[last-'a']++
	}

	return find(&have, target, res, pos)
}

func find(have *[26]int, target string, res []byte, pos int) string {
	for c := int(target[pos]-'a') + 1; c < 26; c++ {
		if have[c] == 0 {
			continue
		}

		res = append(res, byte(c)+'a')
		have[c]--

		for i := 0; i < 26; i++ {
			for have[i] > 0 {
				res = append(res, byte(i)+'a')
				have[i]--
			}
		}

		return string(res)
	}

	if pos == 0 {
		return ""
	}

	pos--

	last := res[len(res)-1]
	res = res[:len(res)-1]
	have[last-'a']++

	return find(have, target, res, pos)
}