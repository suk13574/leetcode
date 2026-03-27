func areSimilar(mat [][]int, k int) bool {
    r, c := len(mat), len(mat[0])
    
    ro := k % c
    for i := 0; i < r; i=i+2 {
        for j := 0; j < c; j++ {
            idx := (j+ro) % c
            if mat[i][j] != mat[i][idx] {
                return false
            }
        }
    }

    for i := 1; i < r; i=i+2 {
        for j := 0; j < c; j++ {
            idx := (j-ro+c) % c
            if mat[i][j] != mat[i][idx] {
                return false
            }
        }
    }

    return true
}