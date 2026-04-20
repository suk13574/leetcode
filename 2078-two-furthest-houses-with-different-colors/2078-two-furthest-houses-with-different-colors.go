func maxDistance(colors []int) int {
    n := len(colors)

    res := 0

    for i := n - 1; i >= 0; i-- {
        if colors[0] != colors[i] {
            res = max(res, i)
            break
        }
    }

    for i := 0; i < n; i++ {
        if colors[i] != colors[n-1] {
            res = max(res, n-1-i)
            break
        }
    }

    return res
}
