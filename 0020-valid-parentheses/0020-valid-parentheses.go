func isValid(s string) bool {
	stack := []byte{}

	pairs := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	for i := 0; i < len(s); i++ {
		ch := s[i]

		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		top := stack[len(stack)-1]

		if top != pairs[ch] {
			return false
		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}