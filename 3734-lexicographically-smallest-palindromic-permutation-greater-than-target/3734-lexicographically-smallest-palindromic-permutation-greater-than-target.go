func lexPalindromicPermutation(s string, target string) string {
	n := len(s)

	cnt := [26]int{}
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
	}

	mid := byte(0)
	odd := 0

	for i := 0; i < 26; i++ {
		if cnt[i]%2 == 1 {
			odd++
			mid = byte('a' + i)
		}

		cnt[i] /= 2
	}

	if odd > 1 {
		return ""
	}

	pos := 0
	half := make([]byte, 0, n/2)

	for pos < n/2 {
		ch := target[pos]

		if cnt[ch-'a'] == 0 {
			break
		}

		half = append(half, ch)
		cnt[ch-'a']--
		pos++
	}

	if pos == n/2 {
		cand := build(half, mid, n)

		if cand > target {
			return cand
		}
	}

	half, ok := find(&cnt, half, target)

	if !ok {
		return ""
	}

	return build(half, mid, n)
}

func find(cnt *[26]int, half []byte, target string) ([]byte, bool) {
	pos := len(half)

	if pos < len(target)/2 {
		for c := int(target[pos]-'a') + 1; c < 26; c++ {
			if cnt[c] == 0 {
				continue
			}

			half = append(half, byte(c+'a'))
			cnt[c]--

			for i := 0; i < 26; i++ {
				for cnt[i] > 0 {
					half = append(half, byte(i+'a'))
					cnt[i]--
				}
			}

			return half, true
		}
	}

	if len(half) == 0 {
		return nil, false
	}

	last := half[len(half)-1]
	half = half[:len(half)-1]
	cnt[last-'a']++

	return find(cnt, half, target)
}

func build(half []byte, mid byte, n int) string {
	res := make([]byte, n)

	for i := 0; i < len(half); i++ {
		res[i] = half[i]
		res[n-1-i] = half[i]
	}

	if n%2 == 1 {
		res[n/2] = mid
	}

	return string(res)
}
