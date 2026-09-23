func minOperations(nums []int, x int) int {
	n := len(nums)

	total := 0
	for _, num := range nums {
		total += num
	}

	S := total - x

    if S < 0 {
        return -1
    }

	if S == 0 {
		return n
	}

	l := 0
    windowSum := 0
    maxLen := -1

	for r := 0; r < n; r++ {
        windowSum += nums[r]

        for windowSum > S {
            windowSum -= nums[l]
            l++
        }

        if windowSum == S {
            maxLen = max(maxLen, r-l+1)
        }
	}

    if maxLen == -1 {
        return -1
    }

    return n - maxLen
}