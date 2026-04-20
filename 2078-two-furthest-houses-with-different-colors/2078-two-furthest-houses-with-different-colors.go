func maxDistance(colors []int) int {
    n := len(colors)

    for i := 0; i < n; i++ {
        r := n - 1 - i
        if colors[0] != colors[r] {
            return r
        }

        l := i
        if colors[l] != colors[n-1] {
            return n - 1 - l
        }
    }

    return 0
}