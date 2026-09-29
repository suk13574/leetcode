func maxSubArray(nums []int) int {
    preMax := nums[0]
    res := preMax

    for i := 1; i < len(nums); i++ {
        cur := max(preMax+nums[i], nums[i])

        res = max(res, cur)
        preMax = cur
    }

    return res
}