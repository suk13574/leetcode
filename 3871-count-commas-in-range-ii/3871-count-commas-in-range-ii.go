func countCommas(n int64) int64 {
    var res int64

    for x := int64(1000); x <= n; x*= 1000 {
        res += n - x + 1
    }

    return res
}