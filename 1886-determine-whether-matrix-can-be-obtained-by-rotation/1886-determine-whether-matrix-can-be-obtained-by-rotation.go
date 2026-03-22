func findRotation(mat [][]int, target [][]int) bool {
	equal := func() bool {
		for i := range mat {
			if !slices.Equal(mat[i], target[i]) {
				return false
			}
		}
		return true
	}

	rotate := func() [][]int {
		n := len(mat)
		rotateMat := make([][]int, n)
		for i := 0; i < n; i++ {
			rotateMat[i] = make([]int, n)
		}

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				rotateMat[j][n-1-i] = mat[i][j]
			}
		}
		return rotateMat
	}

	for k := 0; k < 4; k++ {
		if equal() {
			return true
		}
		mat = rotate()
	}
	return false
}