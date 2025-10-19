func findLexSmallestString(s string, a int, b int) string {
	n := len(s)
	g := gcd(n, b)

	best := s

	for r, cnt := 0, 0; cnt < n/g; r, cnt = (r+b)%n, cnt+1 {
		rot := rotateRight(s, r)

		for addOddCnt := 0; addOddCnt < 10; addOddCnt++ {
			cur := []byte(rot)
			addOdd := (addOddCnt * a) % 10

			for i := 1; i < n; i += 2 {
				cur[i] = addDigit(cur[i], addOdd)
			}

			if b%2 == 1 {
				for addEvenCnt := 0; addEvenCnt < 10; addEvenCnt++ {
					tmp := make([]byte, n)
					copy(tmp, cur)
					addEven := (addEvenCnt * a) % 10
					for i := 0; i < n; i += 2 {
						tmp[i] = addDigit(tmp[i], addEven)
					}
					cand := string(tmp)
					if cand < best {
						best = cand
					}
				}
			} else {
				cand := string(cur)
				if cand < best {
					best = cand
				}
			}
		}
	}

	return best
}

func addDigit(c byte, inc int) byte {
	d := int(c-'0') + inc
	d %= 10
	return byte('0' + d)
}

func rotateRight(s string, r int) string {
	if r == 0 {
		return s
	}
	n := len(s)
	r %= n
	k := (n - r) % n
	return s[k:] + s[:k]
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		return -a
	}
	return a
}