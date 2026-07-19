func findGCD(nums []int) int {
    n := len(nums)
    sort.Ints(nums)
    return gcd(nums[0], nums[n-1])
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }

    return a
}