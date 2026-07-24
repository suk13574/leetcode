func uniqueXorTriplets(nums []int) int {
    n := len(nums)

    if n <= 2 {
        return n
    }

    bits := 0
    x := n

    for x > 0 {
        bits++
        x >>= 1
    }

    return 1 << bits
}