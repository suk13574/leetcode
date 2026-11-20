func missingNumber(nums []int) int {
    n := len(nums)

    for i := 0; i < n; i++ {
        v := nums[i]

        for v != i && v != n {
            nums[v], nums[i] = nums[i], nums[v]
            v = nums[i]
        }
    }

    for i, v := range nums {
        if i != v {
            return i
        }
    }

    return n
}