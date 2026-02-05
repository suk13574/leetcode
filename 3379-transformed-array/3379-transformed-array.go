func constructTransformedArray(nums []int) []int {
    n := len(nums)
    res := make([]int, n)

    for i, num := range nums{
        idx := i
        if num != 0 {
            idx = ((i + num) % n + n) % n
        }

        res[i] = nums[idx]
    }

    return res
}