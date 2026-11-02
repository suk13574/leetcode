func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}

func minMoves2(nums []int) int {
    n := len(nums)

    sort.Ints(nums)

    mid := nums[n/2]

    res := 0
    for i := 0; i < n; i++ {
        res += abs(nums[i]-mid)
    }

    return res
}

