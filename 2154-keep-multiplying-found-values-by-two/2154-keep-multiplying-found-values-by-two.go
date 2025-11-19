import "sort"

func findFinalValue(nums []int, original int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	for _, n := range nums {
		if n == original {
			original *= 2
		}
	}
	return original
}