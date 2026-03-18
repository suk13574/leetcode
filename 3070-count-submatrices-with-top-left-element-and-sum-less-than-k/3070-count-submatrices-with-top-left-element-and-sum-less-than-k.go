func countSubmatrices(grid [][]int, k int) int {
    r, c := len(grid), len(grid[0])

    prefix := make([][]int, r+1)
    for i := 0; i < len(prefix); i++ {
        prefix[i] = make([]int, c+1)
    }

    res := 0
    for i := 0; i < r; i++ {
        for j := 0; j < c; j++ {
            prefix[i+1][j+1] = prefix[i+1][j] + prefix[i][j+1] - prefix[i][j] + grid[i][j]
            if prefix[i+1][j+1] <= k {
                res++
            }
        }
    }

    return res
}