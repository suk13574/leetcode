func reverseSubmatrix(grid [][]int, x int, y int, k int) [][]int {
    for i := 0; i < k/2; i++ {
        top := x + i
        bottom := x + k - i - 1

        for j := 0; j < k; j++ {
            grid[top][y+j], grid[bottom][y+j] = grid[bottom][y+j], grid[top][y+j]
        }

    }
    return grid
}