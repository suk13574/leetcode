func sumAP(n int) int {
	a := 1
	d := 1

	return n * (2*a + (n-1)*d) / 2
}

func zeroFilledSubarray(nums []int) int64 {
	answer := 0
	count := 0
	for _, num := range nums {
		if num == 0 {
			count++
			continue
		}
		if count != 0 {
			answer += sumAP(count)
			count = 0
		}
	}

	if count != 0 {
		answer += sumAP(count)
	}
	return int64(answer)
}
