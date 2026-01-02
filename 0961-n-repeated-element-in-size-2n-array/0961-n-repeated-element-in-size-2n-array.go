import "sort"

func repeatedNTimes(nums []int) int {
	n := len(nums) / 2

	sort.Ints(nums)

	res := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if res == nums[i] {
			count++
		} else {
			res = nums[i]
			count = 1
		}

		if count == n {
			return res
		}
	}

	return 0
}