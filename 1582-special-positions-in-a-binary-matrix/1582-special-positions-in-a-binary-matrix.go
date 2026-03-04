func numSpecial(mat [][]int) int {
    m := len(mat)
    n := len(mat[0])

    rowSum := make([]int, m)
    colSum := make([]int, n)

    for r := 0; r < m; r++ {
        for c := 0; c < n; c++ {
            if mat[r][c] == 1 {
                rowSum[r]++
                colSum[c]++
            }
        }
    }

    res := 0
    for r := 0; r < m; r++ {
        for c := 0; c < n; c++ {
            if mat[r][c] == 1 && rowSum[r] == 1 && colSum[c] == 1 {
                res++
            }
        }
    }

    return res
}