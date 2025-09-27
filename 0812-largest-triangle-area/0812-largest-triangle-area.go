import "math"

func largestTriangleArea(points [][]int) float64 {
	return dfs(points, make([][]int, 0, 3), 0, -1)
}

func dfs(points [][]int, com [][]int, start int, maxArea float64) float64 {
	if len(com) == 3 {
		area := triangleArea(com[0], com[1], com[2])
		return max(maxArea, float64(area))
	}

	for i := start; i < len(points); i++ {
		com = append(com, points[i])
		maxArea = dfs(points, com, i+1, maxArea)
		com = com[:len(com)-1]
	}
	return maxArea
}

func triangleArea(a, b, c []int) float64 {
	area := a[0]*(b[1]-c[1]) + b[0]*(c[1]-a[1]) + c[0]*(a[1]-b[1])
	return math.Abs(float64(area)) / 2.0
}