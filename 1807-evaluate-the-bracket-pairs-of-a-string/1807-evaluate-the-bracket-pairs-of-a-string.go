func evaluate(s string, knowledge [][]string) string {
	m := make(map[string]string)

	for _, kv := range knowledge {
		m[kv[0]] = kv[1]
	}

	idx := 0

	findMap := func() string {
		idx++ // '(' skip
		start := idx

		for s[idx] != ')' {
			idx++
		}

		key := s[start:idx]
		idx++ // ')' skip

		if value, ok := m[key]; ok {
			return value
		}

		return "?"
	}

	res := []byte{}

	for idx < len(s) {
		ch := s[idx]

		if ch == '(' {
			res = append(res, findMap()...)
		} else {
			res = append(res, ch)
			idx++
		}
	}

	return string(res)
}