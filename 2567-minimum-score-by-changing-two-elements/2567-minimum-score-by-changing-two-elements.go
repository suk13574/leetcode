import "sort"

func minimizeSum(nums []int) int {
    n := len(nums)
    if n <=3 {
        return 0
    }

    sort.Ints(nums)

    return min(
        nums[n-1] - nums[2],
        nums[n-3] - nums[0],
        nums[n-2] - nums[1],
    )
}