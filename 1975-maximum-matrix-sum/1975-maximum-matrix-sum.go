func maxMatrixSum(matrix [][]int) int64 {
	n := len(matrix)

	cnt := 0
	minValue := 100001
	var res int64
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			v := matrix[i][j]

			if v < 0 {
				cnt++
				v *= -1
			}

			res += int64(v)
			if v < minValue {
				minValue = v
			}
		}
	}

	if cnt%2 != 0 {
		res -= int64(2 * minValue)
	}

	return res
}