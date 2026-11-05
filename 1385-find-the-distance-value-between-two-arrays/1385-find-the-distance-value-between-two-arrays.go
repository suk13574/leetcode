func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
    res := 0

    for _, v1 := range arr1 {
        distanceValue := true
        for _, v2 := range arr2 {
            if abs(v1-v2) <= d {
                distanceValue = false
                break
            }
        }
        if distanceValue {
            res++
        }
    }

    return res
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}