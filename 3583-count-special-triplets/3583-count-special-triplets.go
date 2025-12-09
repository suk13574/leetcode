func specialTriplets(nums []int) int {
	const MOD int64 = 1000000007

	freqRight := make(map[int]int)
	for _, x := range nums {
		freqRight[x]++
	}

	freqLeft := make(map[int]int)

	var res int64
	for _, x := range nums {
		freqRight[x]--

		target := x * 2

		fl := int64(freqLeft[target])
		fr := int64(freqRight[target])

		cnt := fl * fr % MOD
		res = (res + cnt) % MOD

		freqLeft[x]++
	}
	return int(res)
}