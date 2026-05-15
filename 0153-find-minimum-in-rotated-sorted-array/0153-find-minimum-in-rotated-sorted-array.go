func findMin(nums []int) int {
	n := len(nums)

	if nums[0] < nums[n-1] {
		return nums[0]
	}

	l := 0
	r := n

	for l < r {
		mid := l + (r-l)/2

		if mid+1 < n && nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}

		if nums[l] < nums[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return nums[l]
}
