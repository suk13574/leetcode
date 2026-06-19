func largestAltitude(gain []int) int {
    res := 0
    cur := 0

    for _, g := range gain {
        cur += g
        res = max(res, cur)
    }

    return res
}