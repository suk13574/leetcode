import "sort"

func minSwaps(nums []int) int {
	n := len(nums)

	sorted := make([]int, n)
	copy(sorted, nums)

	sumOfDigits := func(x int) int {
		sum := 0
		for x > 0 {
			sum += x % 10
			x /= 10
		}

		return sum
	}

	sort.Slice(sorted, func(i, j int) bool {
		sum1 := sumOfDigits(sorted[i])
		sum2 := sumOfDigits(sorted[j])

		if sum1 == sum2 {
			return sorted[i] < sorted[j]
		}

		return sum1 < sum2
	})

	targetIdx := make(map[int]int)

	for i, num := range sorted {
		targetIdx[num] = i
	}

	visited := make([]bool, n)
	res := 0

	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}

		cur := i
		cycleSize := 0

		for !visited[cur] {
			visited[cur] = true
			cycleSize++

			val := nums[cur]
			cur = targetIdx[val]
		}

		res += cycleSize - 1
	}

	return res
}