func addBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	carry := byte(0)

	res := make([]byte, 0, max(len(a), len(b))+1)

	for i >= 0 || j >= 0 || carry != 0 {
		sum := carry

		if i >= 0 {
			sum += a[i] - '0'
			i--
		}
		if j >= 0 {
			sum += b[j] - '0'
			j--
		}

		res = append(res, (sum&1)+'0')
		carry = sum >> 1
	}

	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}

	return string(res)
}