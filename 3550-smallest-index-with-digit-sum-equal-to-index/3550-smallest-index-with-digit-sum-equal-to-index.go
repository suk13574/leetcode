func smallestIndex(nums []int) int {
    sumDigits := func(x int) int {
        res := 0
        for x > 0 {
            res += x%10
            x /= 10
        }

        return res
    }

    for i, num := range nums {
        if i == sumDigits(num) {
            return i
        }
    }

    return -1
}