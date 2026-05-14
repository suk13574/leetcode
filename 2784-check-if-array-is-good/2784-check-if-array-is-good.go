import "sort"

func isGood(nums []int) bool {
    sort.Ints(nums)
    n := nums[len(nums)-1]

    if len(nums) != n+1 {
        return false
    }

    for i := 0; i < len(nums)-1; i++ {
        if nums[i] != i+1 {
            return false
        }
    }

    return nums[len(nums)-1] == n
}