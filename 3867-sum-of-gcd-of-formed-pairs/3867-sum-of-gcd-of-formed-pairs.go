func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }

    return a
}

func gcdSum(nums []int) int64 {
    n := len(nums)

    prefixGcd := make([]int, n)
    mx := 0

    for i, num := range nums {
        mx = max(mx, num)
        prefixGcd[i] = gcd(num, mx)
    }

    sort.Ints(prefixGcd)

    var res int64 = 0
    for i := 0; i < n/2; i++ {
        a := prefixGcd[i]
        b := prefixGcd[n-1-i]
        res += int64(gcd(a, b))
    }

    return res
}