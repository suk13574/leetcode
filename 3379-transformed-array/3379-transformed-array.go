func constructTransformedArray(nums []int) []int {
    n := len(nums)
    res := make([]int, n)

    for i := 0; i < n; i++ {
        idx := i
        if nums[i] != 0 {
            idx = (i + nums[i]) % n 
            if idx < 0 {
                idx += n
            }
        }

        res[i] = nums[idx]
    }

    return res
}