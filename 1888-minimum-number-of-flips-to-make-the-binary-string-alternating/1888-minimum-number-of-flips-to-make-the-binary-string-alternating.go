func minFlips(s string) int {
	n := len(s)
	s2 := s + s

	diff0 := 0
	diff1 := 0
	res := n

	for i := 0; i < len(s2); i++ {
		want0 := byte('0')
		want1 := byte('1')

		if i%2 == 1 {
			want0, want1 = want1, want0
		}

		if s2[i] != want0 {
			diff0++
		}
		if s2[i] != want1 {
			diff1++
		}

		if i >= n {
			left := i - n

			leftWant0 := byte('0')
			leftWant1 := byte('1')

			if left%2 == 1 {
				leftWant0, leftWant1 = leftWant1, leftWant0
			}

			if s2[left] != leftWant0 {
				diff0--
			}
			if s2[left] != leftWant1 {
				diff1--
			}
		}

		if i >= n-1 {
			res = min(res, min(diff0, diff1))
		}
	}

	return res
}