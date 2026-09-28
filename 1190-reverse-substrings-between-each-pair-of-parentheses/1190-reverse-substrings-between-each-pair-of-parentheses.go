func reverseParentheses(s string) string {
	n := len(s)

	pair := make([]int, n)
	stack := []int{}

	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			pair[i] = j
			pair[j] = i
		}
	}

	res := []byte{}

	for i, dir := 0, 1; i < n; i += dir {
		if s[i] == '(' || s[i] == ')' {
			i = pair[i]
			dir = -dir
		} else {
			res = append(res, s[i])
		}
	}

	return string(res)
}