func minMoves(nums []int) int {
    n := len(nums)

    maxNum := nums[0]

    for i := 1; i < n; i++ {
        maxNum = max(maxNum, nums[i])
    }

    res := 0

    for _, num := range nums {
        res += maxNum - num
    }

    return res
}