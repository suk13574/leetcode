import "strconv"

func fractionToDecimal(numerator int, denominator int) string {
	var res string
	if (numerator < 0 && denominator > 0) || (numerator > 0 && denominator < 0) {
		res = "-"
	}

	if numerator < 0 {
		numerator *= -1
	}
	if denominator < 0 {
		denominator *= -1
	}

	repeatIdx := -1

	q := numerator / denominator
	r := ""

	division := make(map[int]int)
	for m := numerator % denominator; m != 0; {
		d := (m * 10) / denominator
		m = (m * 10) % denominator

		r += strconv.Itoa(d)

		if idx, ok := division[m]; ok {
			repeatIdx = idx
			break
		}
		idx := len(r)
		division[m] = idx
	}
	fmt.Println(repeatIdx)

	if repeatIdx >= 0 {
		r = r[:repeatIdx] + "(" + r[repeatIdx:] + ")"
	}

	res += strconv.Itoa(q)
	if r != "" {
		res += "." + r
	}

	return res
}