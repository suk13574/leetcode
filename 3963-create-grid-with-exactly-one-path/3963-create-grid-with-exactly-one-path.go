func createGrid(m int, n int) []string {
    grid := make([]string, m)

    for i := 0; i < m; i++ {
        s := []byte{}
        for j := 0; j < n; j++ {
            if i == 0 || j == n-1 {
                s = append(s, '.')
            } else {
                s = append(s, '#')
            }
        }

        grid[i] = string(s)
    }

    return grid
}