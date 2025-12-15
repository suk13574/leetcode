func getDescentPeriods(prices []int) int64 {
    var res int64 = 1
    var descentPeriods int64 = 1
    for i := 1; i < len(prices); i++ {
        if prices[i] == prices[i-1] - 1 {
            descentPeriods++
        } else {
            descentPeriods = 1
        }
        res += descentPeriods
    }

    return res
}