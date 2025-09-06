func count(l, r int) int64 {
	log4 := func(num int) int64 {
		return int64(math.Log(float64(num)) / math.Log(4))
	}

	powerR := log4(r)
	powerL := log4(l)

	outlineLeft := func() float64 {
		fmt.Println(float64(l), math.Pow(4, float64(powerL)))
		return float64(l) - math.Pow(4, float64(powerL))
	}

	outlineRight := func() float64 {
		fmt.Println(math.Pow(4, float64(powerR+1)), float64(r))
		return math.Pow(4, float64(powerR+1)) - float64(r) - 1
	}

	var result int64
	for i := powerL; i <= powerR; i++ {
		cnt := 3 * math.Pow(4, float64(i))
		if i == powerL {
			cnt -= outlineLeft()
		}
		if i == powerR {
			cnt -= outlineRight()
		}

		result += int64(cnt) * (i + 1)

	}

	if result%2 == 0 {
		result /= 2
	} else {
		result /= 2
		result++
	}

	return result
}

func minOperations(queries [][]int) int64 {
	var result int64
	for _, q := range queries {
		l := q[0]
		r := q[1]
		result += count(l, r)
	}

	return result
}