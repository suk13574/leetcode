func minimumCost(nums []int) int {
	res := nums[0]
	newNums := nums[1:]
	sort.Ints(newNums)

	res += newNums[0]
	res += newNums[1]

	return res
}