func countNegatives(grid [][]int) int {
    res := 0
    for _, g := range grid {
        for _, num := range g {
            if num < 0 {
                res++
            }
        }
    }

    return res
}