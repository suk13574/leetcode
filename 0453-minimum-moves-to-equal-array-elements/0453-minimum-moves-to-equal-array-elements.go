func minMoves(nums []int) int {
    n := len(nums)

    minV := nums[0]
    for i := 1; i < n; i++ {
        minV = min(minV, nums[i])
    }

    res := 0
    for i := 0; i < n; i++ {
        res += nums[i] - minV
    }

    return res
}