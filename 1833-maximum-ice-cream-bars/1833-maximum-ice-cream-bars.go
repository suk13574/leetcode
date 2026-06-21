func maxIceCream(costs []int, coins int) int {
	bars := make([]int, 100001)

	for _, cost := range costs {
		bars[cost]++
	}

	res := 0

	for cost := 1; cost < len(bars); cost++ {
		cnt := bars[cost]

		for j := 0; j < cnt; j++ {
			if coins < cost {
				return res
			}

			coins -= cost
			res++
		}
	}

	return res
}