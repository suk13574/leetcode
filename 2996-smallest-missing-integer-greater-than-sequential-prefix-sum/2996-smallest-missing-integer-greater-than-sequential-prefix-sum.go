func missingInteger(nums []int) int {
    sum := nums[0]

    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[i-1]+1 {
            break
        }

        sum += nums[i]
    }

    exists := make(map[int]bool)
    for _, num := range nums {
        exists[num] = true
    }

    for exists[sum] {
        sum++
    }

    return sum
}