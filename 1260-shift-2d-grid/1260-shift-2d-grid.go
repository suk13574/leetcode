func shiftGrid(grid [][]int, k int) [][]int {
    m, n := len(grid), len(grid[0])
    size := m * n

    k %= size

    result := make([][]int, m)
    for i := 0; i < m; i++ {
        result[i] = make([]int, n)

        for j := 0; j < n; j++ {
            newIdx := i*n + j
            oldIdx := (newIdx - k + size) % size

            oldRow := oldIdx / n
            oldCol := oldIdx % n

            result[i][j] = grid[oldRow][oldCol]
        }
    }

    return result
}