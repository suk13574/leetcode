import "sort"

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	n := len(code)

	isValidCode := func(str string) bool {
		if len(str) == 0 {
			return false
		}
		for i := 0; i < len(str); i++ {
			c := str[i]
			if ('0' <= c && c <= '9') ||
				('A' <= c && c <= 'z') ||
				(c == '_') {
				continue
			}
			return false
		}
		return true
	}

	order := []string{"electronics", "grocery", "pharmacy", "restaurant"}
	buckets := map[string][]string{
		"electronics": {},
		"grocery":     {},
		"pharmacy":    {},
		"restaurant":  {},
	}

	for i := 0; i < n; i++ {
		if !isActive[i] || !isValidCode(code[i]) {
			continue
		}

		if _, ok := buckets[businessLine[i]]; !ok {
			continue
		}

		buckets[businessLine[i]] = append(buckets[businessLine[i]], code[i])
	}

	res := make([]string, 0, n)
	for _, bl := range order {
		sort.Strings(buckets[bl])
		res = append(res, buckets[bl]...)
	}
	return res
}