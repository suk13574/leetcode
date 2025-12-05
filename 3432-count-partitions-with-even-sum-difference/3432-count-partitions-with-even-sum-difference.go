func countPartitions(nums []int) int {
    sum := 0
    for _, n := range nums {
        sum += n
    }

    if sum % 2 == 0 {
        return len(nums) - 1
    }
    return 0
}